# Switch Config

[English](#english) | [中文](#中文)

## English

Switch Config is a command-line tool built with Go that allows you to switch your global Git configuration by specifying the path to a Git config file. This tool can be easily installed by using the Go toolchain and `go install` command.

### Installation

To install Switch Config, follow these steps:

1. Make sure that the Go toolchain is installed on your machine.

2. Open a terminal window and run the following command to install the tool:

   ```
   go install github.com/Oasis7311/switch_gitconfig@latest
   ```

3. Once the installation is complete, you can use the tool by typing the following command, followed by the path to the Git config file you wish to use:

   ```
   switch_config -f /path/to/your/gitconfig
   ```

   Note that the path to your Git config file should be an absolute path.

## 中文

Switch Config 是使用 Go 编写的命令行工具，允许您通过指定 Git 配置文件的路径来切换全局 Git 配置。这个工具可以通过使用 Go 工具链和 `go install` 命令轻松安装。

### 安装方法

要安装 Switch Config，请按照以下步骤操作：

1. 确保 Go 工具链已经在您的机器上安装好。

2. 打开终端，然后输入以下命令来安装 Switch Config：

   ```
   go install github.com/Oasis7311/switch_gitconfig@latest
   ```

3. 安装完成后，您可以通过输入以下命令来使用 Switch Config 工具：

   ```
   switch_config -f /path/to/your/gitconfig
   ```

   注意，Git 配置文件的路径应该是绝对路径。