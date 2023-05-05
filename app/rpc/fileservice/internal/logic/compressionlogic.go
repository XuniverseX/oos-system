package logic

import (
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"oos-system/app/rpc/fileservice/internal/svc"
	"oos-system/app/rpc/fileservice/pb"
	"os"
	"path"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompressionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCompressionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompressionLogic {
	return &CompressionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CompressionLogic) Compression(in *pb.CompressionReq) (*pb.CompressionResp, error) {
	// todo: add your logic here and delete this line
	gzipPath := path.Join(l.svcCtx.Config.TempPath, "gzip", in.Hash)
	if _, err := os.Stat(gzipPath); os.IsNotExist(err) {
		os.MkdirAll(gzipPath, os.ModePerm)
	}
	writeFilename := in.Filename
	compressFile(writeFilename, path.Join(gzipPath, in.Hash))
	return &pb.CompressionResp{Success: true}, nil
}

func compressFile(inputPath string, outputPath string) error {

	inputFile, err := os.OpenFile(inputPath, os.O_RDONLY, os.ModePerm)
	fmt.Println(inputPath)
	//outputFile, err := os.Create(outputPath)
	outputFile, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	fmt.Println(outputPath)
	gzipWriter := gzip.NewWriter(outputFile)
	_, err = io.Copy(gzipWriter, inputFile)
	defer gzipWriter.Close()
	defer outputFile.Close()
	defer inputFile.Close()
	if err != nil {
		return err
	}

	return nil
}
