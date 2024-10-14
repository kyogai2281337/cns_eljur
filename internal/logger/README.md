# Logger

## Структура

```go
const (
LTrace LogLevel = iota
LInfo
LWarn
LError
LFatal
)

type LogDest interface {
SetLevel(level LogLevel)
Write(level LogLevel, data string) error // ? maybe i`ll add concatinations
}

type Logger interface {
AddDest(dest ...LogDest)
write(level LogLevel, data string) error
// * Simplifiers:
Trace(data string) error
Info(data string) error
Warn(data string) error
Error(data string) error
Fatal(data string) error
}

```

Для того, чтобы добавить новую абстрактную точку журналирования, необходимо реализовать интерфейс.
На данный момент, стандартные реализации присутствуют для логирования в файл, или в stdout.
Рекомендуется использовать оба варианта за раз, при необходимости, все доп. реализации дописывать в пакет-первоисточник.

Также, в стандартной реализации присутствует структура ***WriteRule***, используящая strings.Builder в качестве формирующего выходное значение звена, что позволило многократно сократить время выполнения, и обьем выделяемой памяти.

Помимо этого, Стандартная реализация также предусматривает timeout в виде значения в миллисекундах вторым параметром при реализации, без возможности исменения(только directly).

Примитивные реализации в поддиректории **/tests**.
