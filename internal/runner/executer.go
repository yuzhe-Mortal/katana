package runner

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/projectdiscovery/gologger"
	"github.com/remeh/sizedwaitgroup"
)

// ExecuteCrawling executes the crawling main loop
func (r *Runner) ExecuteCrawling() error {
	inputs := r.parseInputs()
	if len(inputs) == 0 {
		return errors.New("no input provided for crawling")
	}

	defer r.crawler.Close()

	wg := sizedwaitgroup.New(r.options.Parallelism)
	for _, input := range inputs {
		wg.Add()

		go func(input string) {
			defer wg.Done()

			fmt.Println("start Crawl,input：", input)
			if err := r.crawler.Crawl(input); err != nil {
				gologger.Warning().Msgf("Could not crawl %s: %s", input, err)
			}
		}(input)
	}
	wg.Wait()
	return nil
}
