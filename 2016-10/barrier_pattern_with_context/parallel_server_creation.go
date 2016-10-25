func CreateServers(ctx context.Context, reqs []CreateServerReq) (done []*vmapi.VM, err error) {
        g, gctx := errgroup.WithContext(ctx) // HL

        for _, req := range reqs {
                req := req // https://golang.org/doc/faq#closures_and_goroutines
                g.Go(func() error { // HL
                        srv, err := CreateServer(gctx, &req)
                        if err == nil {
                                done = append(done, srv)
                        }
                        return err
                })
        }

        if err := g.Wait(); err != nil {
                logger.Errorf("create server failed: %s", err)
                for _, srv := range done {
                        _ = DeleteServer(ctx, srv.Name) // remove completed ones // HL
                }
                return nil, err
        }
        return done, nil
}
