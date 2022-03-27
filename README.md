# gufetch
gufetch is a CLI tool written in Go that can fetch your GitHub or Gitlab profile

## Setting up on Arch Linux
<!--### AUR 
```bash
yay -S gufetch
```-->
### Manually
```bash
git clone https://github.com/shanepaton/gufetch.git
cd gufetch
makepkg -si
```
## Configuration
- Upon installation a configuration file will be created in `~/.config/gufetch/config.yaml`
```yaml
token:
    gitlab: "gitlab token"
defualt:
    githubUsername: "github username"
    gitlabID: "gitlab id"
```
The defualt names are used when the program is called without a specified username/id.<br>
To fetch Gitlab you need to provide and access token in the config file.


gufetch is created by Shane Paton under the MIT Licence
