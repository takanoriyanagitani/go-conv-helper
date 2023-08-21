import logging

import grpc
import conv_pb2
import conv_pb2_grpc

import colorsys

rgb2hsv = colorsys.rgb_to_hsv
hsv2rgb = colorsys.hsv_to_rgb

addr = "localhost:2151"

def main():
	logging.basicConfig()

	with grpc.insecure_channel(addr) as chan:
		stub = conv_pb2_grpc.HsvCmdRecvServiceStub(chan)
		res = stub.GetCmd(
			conv_pb2.HsvOrder.GetCmdRequest(
				request_id=conv_pb2.Uuid(
					hi=0xffffffff,
					lo=0xcafef00d,
				),
			),
		)
		s2 = conv_pb2_grpc.HsvEvtServiceStub(chan)
		r2 = s2.Converted(
			conv_pb2.HsvEvt.ConvertedRequest(
				meta = conv_pb2.Meta(
					request_id=conv_pb2.Uuid(
						hi=0xffff,
						lo=0xf00d,
					),
					reply_id=conv_pb2.Uuid(
						hi=0xffff,
						lo=0xcafe,
					),
				),
				converted = conv_pb2.Rgb(
					red = 0.0,
					green = 1.0,
					blue = 0.0,
				),
			),
		)
		pass
	pass

main()
