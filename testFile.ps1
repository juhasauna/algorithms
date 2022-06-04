$testFileName = [IO.Path]::ChangeExtension($args[0], [NullString]::Value)
if (!$testFileName) {
    echo "You forgot to give the commandline argument for testFileName. I QUIT!"
    Return
}
$cylanceWhiteListedFolder = "C:\MO\Motors_Generators\tech\SDC\"
$exeName = [System.IO.Path]::GetFileNameWithoutExtension($testFileName) + ".exe"
$exeFullPath = $cylanceWhiteListedFolder + $exeName

$wd = $pwd.ToString()
$projectName = Split-Path $wd -Leaf     

$packageName = $testFileName -replace [RegEx]::Escape($wd), "" | Split-Path -Parent
$projectPackageName = $projectName + $packageName

go test $projectPackageName -c -o $exeFullPath; 
& $exeFullPath "test" $testFileName 2>&1 | ForEach-Object{ "$_" }
    