# Requires Windows 10/11 with winget available
$ErrorActionPreference = "Stop"

# Docker Desktop
winget install --id Docker.DockerDesktop -e --source winget

# Go
$goVersion = "1.23.3"
winget install --id GoLang.Go --version $goVersion -e --source winget

# Node/npm
$nodeVersion = "20"
winget install OpenJS.NodeJS.LTS --version $nodeVersion -e --source winget

Write-Host "Done. Log out/in after Docker Desktop install. Start Docker Desktop once to finish setup."