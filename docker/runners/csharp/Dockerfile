FROM mcr.microsoft.com/dotnet/sdk:8.0-alpine

WORKDIR /sandbox
COPY runner.sh /runner.sh
RUN chmod +x /runner.sh

# Pre-warm NuGet cache for a simple console app
RUN dotnet new console -o /template_cache && rm -rf /template_cache

ENTRYPOINT ["/runner.sh"]
