package browser

import (
	"context"
	"github.com/chromedp/chromedp"
	"net/url"
	"time"
)

type Service struct {
	Context context.Context
	Cancel  context.CancelFunc
}

func (s *Service) Init() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoFirstRun,
		chromedp.NoSandbox,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Headless,
		chromedp.DisableGPU,
		chromedp.WindowSize(1280, 720),
	)
	s.Context, s.Cancel = chromedp.NewExecAllocator(context.Background(), opts...)
}

func (s *Service) Screenshot(el string, url *url.URL) ([]byte, error) {
	var buf []byte

	pageCtx, cancel := chromedp.NewContext(s.Context)
	defer cancel()

	if err := chromedp.Run(pageCtx, chromedp.Navigate(url.String())); err != nil {
		return nil, err
	}

	taskContext, cancel := context.WithTimeout(pageCtx, 30*time.Second)
	defer cancel()

	task := chromedp.Tasks{
		chromedp.Navigate(url.String()),
		chromedp.Sleep(1 * time.Second),
		chromedp.Screenshot(el, &buf, chromedp.NodeVisible),
	}

	if err := chromedp.Run(taskContext, task); err != nil {
		return nil, err
	}

	return buf, nil
}

func (s *Service) Close() {
	s.Cancel()
}
