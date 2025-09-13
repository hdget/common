package meta

import (
	"context"
)

type Meta interface {
	//GetAppId(ctx context.Context) string
	//GetClient(ctx context.Context) string
	//GetRelease(ctx context.Context) string
	//GetUid(ctx context.Context) int64  // 获取用户的id
	//GetTsn(ctx context.Context) string // 获取租户的sn
	//GetTid(ctx context.Context) int64  // 获取租户的id
	//GetRoleIds(ctx context.Context) []int64
	//GetCaller(ctx context.Context) string
	// DEPRECATED
	OldGetRoles(ctx context.Context) []*Role
	OldGetRoleValues(ctx context.Context) []string
	OldGetRoleIds(ctx context.Context) []int64
	OldGetPermIds(ctx context.Context) []int64
}

//	func (m grpcMetaReaderImpl) GetAppId(ctx servicectx.Context) string {
//		return FromContext(ctx).GetString(KeyAppId)
//	}
//
//	func (m grpcMetaReaderImpl) GetClient(ctx servicectx.Context) string {
//		return getGrpcMdFirstValue(ctx, KeyClient)
//	}
//
//	func (m grpcMetaReaderImpl) GetRelease(ctx servicectx.Context) string {
//		return getGrpcMdFirstValue(ctx, KeyRelease)
//	}
//
//	func (m grpcMetaReaderImpl) GetCaller(ctx servicectx.Context) string {
//		return getGrpcMdFirstValue(ctx, KeyCaller)
//	}
//
//	func (m grpcMetaReaderImpl) GetRoleIds(ctx servicectx.Context) []int64 {
//		return convert.CsvToInt64s(getGrpcMdFirstValue(ctx, KeyRoleIds))
//	}
//
//	func (m grpcMetaReaderImpl) GetUid(ctx servicectx.Context) int64 {
//		return cast.ToInt64(getGrpcMdFirstValue(ctx, KeyUid))
//	}
//
//	func (m grpcMetaReaderImpl) GetTsn(ctx servicectx.Context) string {
//		return getGrpcMdFirstValue(ctx, KeyTsn)
//	}
//
//	func (m grpcMetaReaderImpl) GetTid(ctx servicectx.Context) int64 {
//		return cast.ToInt64(getGrpcMdFirstValue(ctx, KeyTid))
//	}
//
