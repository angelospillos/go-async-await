package async

type Task func() (interface{}, error)

func RunAsync(tasks ...Task) ([]interface{}, []error) {
    var wg sync.WaitGroup
    results := make([]interface{}, len(tasks))
    errors := make([]error, len(tasks))

    for i, task := range tasks {
        wg.Add(1)
        go func(i int, t Task) {
            defer wg.Done()
            result, err := t()
            results[i] = result
            errors[i] = err
        }(i, task)
    }

    wg.Wait()
    return results, errors
}

func RunSingleAsync(task Task) (interface{}, error) {
    resultChan := make(chan interface{})
    errorChan := make(chan error)

    go func() {
        result, err := task()
        resultChan <- result
        errorChan <- err
    }()

    return <-resultChan, <-errorChan
}
