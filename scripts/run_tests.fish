#!/usr/bin/env fish

# * Note, only meant to be run from the root dir.

set packages pkg/api pkg/messages pkg/utils

for dir in $packages
    # Only proceed if it's a dir.
    if test -d $dir
        # Change to each test's package dir.
        cd $dir

        # Run all tests.
        echo "Running tests in: $dir"
        go test -v

        if not test $status -eq 0
            echo "Tests failed in: $dir"
            exit $status
        end

        # Go back to the root directory.
        cd ../..
    end
end
