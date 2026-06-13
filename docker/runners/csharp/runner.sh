#!/bin/sh

# Expects Solution.cs to be mounted at /sandbox/Program.cs
# We generate a project on the fly if it doesn't exist to run Program.cs
if [ ! -f "Solution.csproj" ]; then
    dotnet new console --force > /dev/null
fi

cp /sandbox/Solution.cs /sandbox/Program.cs

if ! dotnet build -c Release > /dev/null; then
    echo "Compilation Error"
    dotnet build -c Release
    exit 1
fi

dotnet run -c Release --no-build
