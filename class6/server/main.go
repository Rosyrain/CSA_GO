package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	fieldmask_utils "github.com/mennanov/fieldmask-utils"
	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb" // 导入 Empty 消息类型包
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"log"
	"net"
	"server/mysql"
	"server/pb"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) CreateUser(ctx context.Context, in *pb.CreateUserInfo) (*pb.UserInfoResponse, error) {
	nickname := in.GetNickName()
	password := in.GetPassWord()
	mobile := in.GetMobile()

	if err := mysql.CreateUser(nickname, password, mobile); err != nil {
		log.Fatalf("mysql.CreateUser failed,err:%v", err)
	}

	user, err := mysql.GetUserInfo(nickname)
	if err != nil {
		log.Fatalf("mysql.GetUserInfo failed,err:%v", err)
	}

	return &pb.UserInfoResponse{
		Id:       user.ID,
		NickName: user.Nickname,
		PassWord: user.Password,
		BirthDay: user.BirthDay,
		Gender:   user.Gender,
	}, nil
}

func (s *server) GetUserById(ctx context.Context, in *pb.IdRequest) (*pb.UserInfoResponse, error) {
	id := in.GetId()

	user, err := mysql.GetUserByID(id)
	if err != nil {
		log.Fatalf("mysql.GetUserInfo failed,err:%v", err)
	}

	return &pb.UserInfoResponse{
		Id:       user.ID,
		NickName: user.Nickname,
		PassWord: user.Password,
		BirthDay: user.BirthDay,
		Gender:   user.Gender,
	}, nil

}

func (s *server) UpdateUser(ctx context.Context, in *pb.UpdateUserInfo) (*emptypb.Empty, error) {
	paths := []string{"user.id", "user.nickname", "user.brithDay", "user.gender"}

	req := pb.UpdateUserRequest{
		Op:   "rosyrain",
		User: in,

		UpdateMask: &fieldmaskpb.FieldMask{Paths: paths},
	}

	//Server
	mask, _ := fieldmask_utils.MaskFromProtoFieldMask(req.UpdateMask, generator.CamelCase)
	var userDst = make(map[string]interface{})
	// 将数据读取到map[string]interface{}
	// fieldmask-utils支持读取到结构体等，更多用法可查看文档。
	fieldmask_utils.StructToMap(mask, req.User, userDst)
	// do update with bookDst
	fmt.Printf("bookDst:%#v\n", userDst)
	err := mysql.UpdateUserInfo(req.User.Id, req.User.NickName, req.User.BirthDay, req.User.Gender)
	return nil, err
}

func (s *server) CheckPassWord(ctx context.Context, in *pb.PasswordCheckInfo) (*pb.CheckResponse, error) {
	passsword := in.GetPassword()
	enctryptedPassword := in.GetEncryptedPassword()
	success, err := mysql.CheckPassWord(passsword, enctryptedPassword)
	return &pb.CheckResponse{Success: success}, err
}

func main() {
	//	初始化数据库
	if err := mysql.Init(); err != nil {
		log.Fatalf("mysql.Init failed,err:%v", err)
	}
	defer mysql.Close()

	//	监听端口
	l, err := net.Listen("tcp", "8989")
	if err != nil {
		log.Fatalf("net.Listen faild,err:%v", err)
	}

	//创建grpc服务
	s := grpc.NewServer()
	//注册服务
	pb.RegisterUserServiceServer(s, &server{})

	//启动服务
	err = s.Serve(l)
	if err != nil {
		log.Fatalf("s.Server faild,err:%v", err)
	}

}
