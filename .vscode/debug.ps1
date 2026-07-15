param(
  [Parameter(Mandatory=$true)]
  [string]$goFile
)

$base = [System.IO.Path]::GetFileName($goFile)
if ($base -notmatch '_test\.go$') {
  Write-Error "Not a *_test.go file: $base"
  exit 1
}

$prefix = $base -replace '_test\.go$', ''
$testName = "Test_$prefix"

$pkgDir = Split-Path -Parent $goFile
$testExe = Join-Path $pkgDir "__debug_test.exe"

Push-Location $pkgDir
try {
  go test -c -o $testExe .
  if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

  dlv exec $testExe -- -test.run ("^" + [Regex]::Escape($testName) + "$") -test.v
} finally {
  if (Test-Path $testExe) {
    Remove-Item $testExe -Force
  }
  Pop-Location
}