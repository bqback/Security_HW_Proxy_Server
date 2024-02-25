# Generate local CA
Write-Output "Generating certificate authority"
cmd  /c "bash.exe ./scripts/gen_ca.sh"
if($?) {
    Write-Output "Done"
    
    # Add local CA
    Write-Output "Adding certificate authority to current user's CA store"
    Import-Certificate -FilePath 'proxy-serv-ca.crt' -CertStoreLocation Cert:\CurrentUser\Root
    Write-Output "Done"
} else {
    Write-Output $LASTEXITCODE
    Write-Error "Failed to run gen script"
}
