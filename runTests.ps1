$wd = $pwd.ToString()
$packageName = "." + ($args[0] | Split-Path -Parent).ToLower().Replace($wd.ToLower(),"") + "\"

$testFileName1 = ($args[0] | Split-Path -Leaf).TrimEnd(".go")
$testFileName2 = $testFileName1.TrimEnd("_test")
$isVerbose = $testFileName1 -eq $testFileName2
$testFileName2 = "Test_" + $testFileName2

$cylanceWhiteListedFolder = "C:\MO\Motors_Generators\tech\SDC\"
$exeFullPath = $cylanceWhiteListedFolder + "testingGoProgram2.exe"

if ($isVerbose) {
    go test -v -run $testFileName2 $packageName -o $exeFullPath
} else {
    go test -run $testFileName2 $packageName -o $exeFullPath
}
