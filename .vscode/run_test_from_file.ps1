# -count=1 avoids cached results.
# Regex::Escape handles any odd characters safely (usually not needed, but harmless).

param(
  [Parameter(Mandatory=$true)]
  [string]$goFile
)

# Ensure it is a *_test.go file
$base = [System.IO.Path]::GetFileName($goFile)
if ($base -notmatch '_test\.go$') {
  Write-Error "Not a *_test.go file: $base"
  exit 1
}

# bom_test.go -> bom
$prefix = $base -replace '_test\.go$', ''
$testName = "Test_$prefix"

# Run in the package directory of the file
$pkgDir = Split-Path -Parent $goFile
Push-Location $pkgDir
try {
  # Run exactly that test
  go test -run ("^" + [Regex]::Escape($testName) + "$") -count=1 -v .
} finally {
  Pop-Location
}