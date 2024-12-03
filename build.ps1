# Go Cross-Platform Build Script

# Configuration
$projectName = "gcms"
$mainPackagePath = "."  # Path to your main.go file
$outputDir = ".\dist"

# Create output directory if it doesn't exist
if (!(Test-Path -Path $outputDir)) {
    New-Item -ItemType Directory -Force -Path $outputDir
}

# Platforms to build for
$platforms = @(
    @{OS="windows"; Arch="amd64"; Ext=".exe"},
    @{OS="windows"; Arch="386"; Ext=".exe"},
    @{OS="darwin"; Arch="amd64"; Ext=""},
    @{OS="darwin"; Arch="arm64"; Ext=""},
    @{OS="linux"; Arch="amd64"; Ext=""},
    @{OS="linux"; Arch="386"; Ext=""},
    @{OS="linux"; Arch="arm"; Ext=""},
    @{OS="linux"; Arch="arm64"; Ext=""}
)

# Build function
function Build-GoProject {
    param (
        [string]$OS,
        [string]$Arch,
        [string]$Ext
    )

    $env:GOOS = $OS
    $env:GOARCH = $Arch

    $outputName = "$projectName-$OS-$Arch$Ext"
    $outputPath = Join-Path $outputDir $outputName

    Write-Host "Building for $OS/$Arch..."
    
    go build -o $outputPath $mainPackagePath

    if ($LASTEXITCODE -eq 0) {
        Write-Host "Successfully built $outputName" -ForegroundColor Green
    } else {
        Write-Host "Failed to build $outputName" -ForegroundColor Red
    }
}

# Perform builds
foreach ($platform in $platforms) {
    Build-GoProject -OS $platform.OS -Arch $platform.Arch -Ext $platform.Ext
}

# Optional: Create checksums
Write-Host "Generating checksums..." -ForegroundColor Cyan
Get-ChildItem $outputDir | ForEach-Object {
    $hash = (Get-FileHash $_.FullName -Algorithm SHA256).Hash
    $hashFile = $_.FullName + ".sha256"
    $hash | Out-File $hashFile
    Write-Host "Checksum for $($_.Name): $hash"
}

Write-Host "Build process completed. Binaries are in $outputDir" -ForegroundColor Green