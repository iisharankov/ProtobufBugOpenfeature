# Golang FeatureFlag provider
This Golang provider is designed to talk to the a [flag-managment-system](REDACTED) (FMS), which in turn communicates to a FeatureFlag database containingflag metadata.  

## Run

To run this, simply call
```
go run
./openfeature-provider-golang
``` 

If you get an error saying 
```
provider/provider.go:7:2: no required module provides package REDACTED/evaluation/evaluation/v1; to add it:
        go get REDACTED/evaluation/evaluation/v1
```
make sure to follow the [import the featureflag-go-local package](#importing-golang-package-from-artifactory) section.

# Importing featureflag-go-local package from Artifactory

For this to work locally, you'll need to have the `featureflag-go-local` module setup with golang. Follow the below steps to do this:

1. export GOPROXY="<https://USERNAME:PASS@artifactory.REDACTED.com/artifactory/api/go/featureflag-go-local>"

   - Get a token, don't pass your password.
   - Make sure user being used has at the very least read access.
2. `export GONOSUMDB=REDACTED.com` 

   - This tells Golang not to verify modules from REDACTED. See the extra info/notes. May not be necessary if checksum already published/

3. `go get REDACTED.com/evaluation`

   - Should see:  
   `go: downloading REDACTED.com/evaluation v1.0.0 `  
   `go: added REDACTED.com/evaluation v1.0.0`
4. Check `$GOPATH/pkg/mod/` for the package. It should be under the new `REDACTED.com` directory. If you haven't set a custom `GOPATH`, the default location is usually `$HOME/go/pkg/mod/`
5. Navigate to the correct package in the module you would like to import. Check it's there. In this case, `REDACTED.com/evaluation@v1.0.0/evaluation/v1` should exist
6. Import into your module/package as needed. In this case, the correct path is `v1 "REDACTED.com/evaluation/evaluation/v1"` 

