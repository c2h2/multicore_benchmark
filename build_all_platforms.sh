#!/usr/bin/env bash

package=multicore_bench.go
package_name=multicore_bench
platforms=("linux/amd64" "linux/386" "linux/arm" "linux/arm64" "linux/mips" "linux/mips64" "darwin/arm64" "darwin/amd64" "freebsd/amd64" "freebsd/386" "windows/386" "windows/amd64")
rm -rf build
mkdir -p build/tarballs

for platform in "${platforms[@]}"
do	
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name=$package_name'-'$GOOS'-'$GOARCH
	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi	

	env GOOS=$GOOS GOARCH=$GOARCH go build -o build/$output_name $package
	tar cvfz build/tarballs/$output_name.tar.gz build/$output_name
	if [ $? -ne 0 ]; then
   		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	fi
done
