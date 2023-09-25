package main

import (
	"context"
	"flag"
	"io"
	"os"

	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/frontend/gateway/grpcclient"
	"github.com/moby/buildkit/util/appcontext"
	"github.com/moby/buildkit/util/bklog"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/sozercan/aikit/pkg/aikit/config"
	"github.com/sozercan/aikit/pkg/aikit2llb"
	"github.com/sozercan/aikit/pkg/build"
	"google.golang.org/grpc/grpclog"
)

var (
	graph    bool
	filename string
)

func main() {
	bklog.L.Logger.SetOutput(os.Stderr)
	grpclog.SetLoggerV2(grpclog.NewLoggerV2WithVerbosity(bklog.L.WriterLevel(logrus.InfoLevel), bklog.L.WriterLevel(logrus.WarnLevel), bklog.L.WriterLevel(logrus.ErrorLevel), 1))

	flag.BoolVar(&graph, "graph", false, "output a graph and exit")
	flag.StringVar(&filename, "filename", "aikitfile.yaml", "the file to read from")
	flag.Parse()

	ctx := appcontext.Context()

	if graph {
		if err := printGraph(ctx, filename, os.Stdout); err != nil {
			bklog.L.WithError(err).Fatal("error printing graph")
			os.Exit(1)
		}
		os.Exit(0)
	}

	if err := grpcclient.RunFromEnvironment(ctx, build.Build); err != nil {
		bklog.L.WithError(err).Fatal("error running frontend")
		os.Exit(137)
	}
}

func printGraph(ctx context.Context, filename string, out io.Writer) error {
	c, err := config.NewFromFilename(filename)
	if err != nil {
		return errors.Wrap(err, "opening config file")
	}
	st, _ := aikit2llb.Aikit2LLB(c)
	dt, err := st.Marshal(ctx)
	if err != nil {
		return errors.Wrap(err, "marshaling llb state")
	}

	return llb.WriteTo(dt, out)
}
