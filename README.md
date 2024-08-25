![](https://socialify.git.ci/testzboy/SudoSnatcher/image?description=1&descriptionEditable=After%20obtaining%20regular%20or%20privileged%20user%20login%20access%2C%20this%20tool%20can%20be%20used%20to%20capture%20sudo%20passwords.&font=Inter&forks=1&issues=1&language=1&logo=https%3A%2F%2Favatars.githubusercontent.com%2Fu%2F178143062%3Fs%3D96%26v%3D4&name=1&owner=1&stargazers=1&theme=Auto)

# Linux 后渗透工具：SudoSnatcher

[English](https://github.com/testzboy/SudoSnatcher/blob/main/README_EN.md) | [Chinese](https://github.com/testzboy/SudoSnatcher/blob/main/README.md)

# 免责说明

此工具仅限于安全研究和教学，作者不承担任何法律和相关责任，密码仅保存在本地，不提供云上传功能。

# 使用场景

用于在获取普通用户/特权用户登录权限后，进行sudo密码抓取。用户可使用此程序制造水坑，窃取合法用户的密码。

# 特点

- 用户无感
- 无需依赖
- 自动清理痕迹

# 使用方法

[Releases](https://github.com/testzboy/SudoSnatcher/releases) 中下载指定版本的二进制文件，可通过默认配置直接运行或指定参数运行。

参数`i`指定后渗透路径，参数`o`指定密码路径。

默认保存的密码本为`/tmp/.pass`：

```
$ ./linux_amd64_SudoSnatcher -h
Usage of ./linux_amd64_SudoSnatcher:
  -i string
    	Path to the script for the alias (default "/tmp/.cache")
  -o string
    	Output file path for saved passwords (default "/tmp/.pass")
```

运行后输入`quit`则自动清理痕迹，恢复设备默认状态，仅保留生成的密码本。

# 密码类型

密码分为三种状态：

```
test:111111:fail
test:000000:success
test:000000:valid
```

fail：错误密码

success：正确密码

valid：sudo session 环境下的待验证密码



![image](https://github.com/user-attachments/assets/8f171a3b-2717-44e2-9348-0eb0abbe4017)
