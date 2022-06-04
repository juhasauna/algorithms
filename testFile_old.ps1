$testFileName = $args[0]
if (!$testFileName) {
    echo "You forgot to give the commandline argument for testFileName. I QUIT!"
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
$exeFullPath = $cylanceWhiteListedFolder + $testFileName + ".exe"

go test -c -o $exeFullPath; 
& $exeFullPath "test" $testFileName 2>&1 | ForEach-Object{ "$_" }
cd $wd
    