$testFunctionName = $args[0]
if (!$testFunctionName) {
    echo "You forgot to give the commandline argument for testFunctionName. I QUIT!"
    Return
}
$testPath = $args[1]
$wd = $pwd.ToString()
if (!$testPath) {
    echo "no test path provided"
} else {
    cd $testPath
}
$cylanceWhiteListedFolder = "C:\MO\Motors_Generators\tech\SDC\"
$exeFullPath = $cylanceWhiteListedFolder + "algorithms.exe"
# if (Test-Path $exeFullPath) {
#     rm $exeFullPath
# }
go test -c -o $exeFullPath; 
& $exeFullPath "-test.run" $testFunctionName 2>&1 | ForEach-Object{ "$_" }
cd $wd
    