@echo off
setlocal

pushd %~dp0

for /f "delims=" %%x in (version) do set VERSION=%%x

docker build  --progress plain --build-arg VERSION=%VERSION% -t dcjulian29/bind:%VERSION% .

if %errorlevel% neq 0 GOTO FINAL

docker tag dcjulian29/bind:%VERSION% dcjulian29/bind:latest

:FINAL

popd

set IMAGE_VERSION=%VERSION%

goreleaser --snapshot --skip-publish --clean

endlocal
