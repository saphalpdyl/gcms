# Go Cross-Platform Build Script
$projectName = "gcms"
$mainPackagePath = "."
$outputDir = ".\dist"

if (!(Test-Path -Path $outputDir)) {
    New-Item -ItemType Directory -Force -Path $outputDir | Out-Null
}

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

# Parallel build function
function Build-GoProjectParallel {
    param (
        [string]$OS,
        [string]$Arch,
        [string]$Ext,
        [string]$ProjectName,
        [string]$MainPackagePath,
        [string]$OutputDir
    )

    $job = Start-Job -ScriptBlock {
        param($OS, $Arch, $Ext, $ProjectName, $MainPackagePath, $OutputDir)
        
        $env:GOOS = $OS
        $env:GOARCH = $Arch

        $outputName = "$ProjectName-$OS-$Arch$Ext"
        $outputPath = Join-Path $OutputDir $outputName

        $result = @{
            OS = $OS
            Arch = $Arch
            Success = $false
            OutputPath = $outputPath
        }

        try {
            go build -o $outputPath $MainPackagePath
            $result.Success = $true
        }
        catch {
            $result.Success = $false
        }

        return $result
    } -ArgumentList $OS, $Arch, $Ext, $ProjectName, $MainPackagePath, $OutputDir
    
    return $job
}

# Start parallel builds
$jobs = $platforms | ForEach-Object {
    Build-GoProjectParallel -OS $_.OS -Arch $_.Arch -Ext $_.Ext `
                            -ProjectName $projectName `
                            -MainPackagePath $mainPackagePath `
                            -OutputDir $outputDir
}

# Wait for all jobs and process results
$results = $jobs | Wait-Job | Receive-Job

# Display results
foreach ($result in $results) {
    if ($result.Success) {
        Write-Host "Successfully built $($result.OS)/$($result.Arch)" -ForegroundColor Green
    } else {
        Write-Host "Failed to build $($result.OS)/$($result.Arch)" -ForegroundColor Red
    }
}

# Clean up jobs
$jobs | Remove-Job

# Optional: Create checksums
Write-Host "Generating checksums..." -ForegroundColor Cyan
Get-ChildItem $outputDir | ForEach-Object {
    $hash = (Get-FileHash $_.FullName -Algorithm SHA256).Hash
    $hashFile = $_.FullName + ".sha256"
    $hash | Out-File $hashFile
    Write-Host "Checksum for $($_.Name): $hash"
}

Write-Host "Parallel build process completed. Binaries are in $outputDir" -ForegroundColor Green