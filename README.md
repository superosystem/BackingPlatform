# Sport Shop Net

## Usage

## Development

* SportStore

```bash
# Command to create project
dotnet new globaljson --sdk-version 6.0.100 --output src/SportsStore
dotnet new web --no-https --output src/SportsStore --framework net6.0
dotnet new sln -o SportShopNet
cd SportShopNet/SportShopNet.sln ../
dotnet sln SportShopNet.sln add src/SportsStore

# Command to create test project
dotnet new xunit -o tests/SportsStore.Tests --framework net6.0
dotnet sln SportShopNet.sln add tests/SportsStore.Tests
dotnet add tests/SportsStore.Tests reference src/SportsStore

# Command to use Moq package to create mock objects
dotnet add tests/SportsStore.Tests package Moq --version 4.18.2

```
