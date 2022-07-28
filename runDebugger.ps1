$packageName = "algorithms/numberTheory"
$testName = "Test_primes"

$cylanceWhiteListedFolder = "C:\MO\Motors_Generators\tech\SDC\"
$exeFullPath = $cylanceWhiteListedFolder + "goDebugger.exe"

go test -v -run $testName $packageName -o $exeFullPath