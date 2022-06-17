$wd = $pwd.ToString()
$packageName = "." + ($args[0] | Split-Path -Parent).ToLower().Replace($wd.ToLower(),"") + "\"

$testName = ($args[0] | Split-Path -Leaf).TrimEnd(".go").TrimEnd("_test")
$testName = "Test_" + $testName

$cylanceWhiteListedFolder = "C:\MO\Motors_Generators\tech\SDC\"
$exeFullPath = $cylanceWhiteListedFolder + "testingGoProgram2.exe"

go test -v -run $testName $packageName -o $exeFullPath