try
{
    go test -bench BenchmarkSolverSize15 . -trace="./.tracedata/trace.out" -cpuprofile "./.tracedata/cpu.prof" -memprofile "./.tracedata/mem.prof"
    $cpu = Start-Process "go.exe" "tool","pprof","-http",":","./.tracedata/cpu.prof" -PassThru
    $mem = Start-Process "go.exe" "tool","pprof","-http",":","./.tracedata/mem.prof" -PassThru
    $trc = Start-Process "go.exe" "tool","trace","./.tracedata/trace.out" -PassThru
    Write-Output "Waiting for Ctrl+C to exit"
    while($true) {
        Start-Sleep -Seconds 60
    }
}
finally
{
    Stop-Process $cpu
    Stop-Process $mem
    Stop-Process $trc

    Stop-Process -Name "pprof"
    Stop-Process -Name "trace"
}