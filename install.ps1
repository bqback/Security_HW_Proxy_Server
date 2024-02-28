$MIGRATOR_PASSWORD=$args[0]
$POSTGRES_USER=$args[1]
$POSTGRES_PASSWORD=$args[2]
$POSTGRES_DB=$args[3]
$POSTGRES_HOST=$args[4]

if(!(Test-Path -Path "config/.env")){
    Write-Output ".env file not found, creating"
    New-Item -Path "config" -Name ".env" -ItemType File 
    Add-Content -Path "config/.env" "MIGRATOR_PASSWORD=`"$MIGRATOR_PASSWORD`""
    Add-Content -Path "config/.env" "POSTGRES_USER=`"$POSTGRES_USER`""
    Add-Content -Path "config/.env" "POSTGRES_PASSWORD=`"$POSTGRES_PASSWORD`""
    Add-Content -Path "config/.env" "POSTGRES_DB=`"$POSTGRES_DB`""
    Add-Content -Path "config/.env" "POSTGRES_HOST=`"$POSTGRES_HOST`""
} else {
    Write-Output ".env file found"
}

# Check if local CA file doesn't exist
if(!(Test-Path -Path "proxy-serv-ca.crt")){
    # Generate local CA
    Write-Output "Generating certificate authority"
    cmd  /c "bash.exe ./scripts/gen_ca.sh"
    if($?) {
        Write-Output $LASTEXITCODE
        Write-Error "Failed to run gen script"
        Exit
    }
    Write-Output "Done"
} else {
    Write-Output "CA found"
}

# Add local CA
Write-Output "Adding certificate authority to current user's CA store"
Import-Certificate -FilePath 'proxy-serv-ca.crt' -CertStoreLocation Cert:\CurrentUser\Root
if($?) {
    Write-Output "Done"
} else {
    Write-Output $LASTEXITCODE
    Write-Error "Failed to add certificate"
    Exit
}