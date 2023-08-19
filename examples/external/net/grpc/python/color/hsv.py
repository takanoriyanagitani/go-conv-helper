import colorsys
import logging

from concurrent import futures

import grpc
import conv_pb2
import conv_pb2_grpc

rgb2hsv = colorsys.rgb_to_hsv
hsv2rgb = colorsys.hsv_to_rgb

def rgb2reply(r,g,b):
	return conv_pb2.FromHsvResponse(
		rgb=conv_pb2.Rgb(
			red=r,
			green=g,
			blue=b,
		),
	)

class ColorConvSvc(conv_pb2_grpc.ColorConvService):
	def FromHsv(self, request, context):
		hsv = request.hsv
		h = hsv.hue
		s = hsv.saturation
		v = hsv.value
		rgb = hsv2rgb(h, s, v)
		return rgb2reply(*rgb)
	pass

def main():
	logging.basicConfig()
	server = grpc.server(futures.ThreadPoolExecutor(max_workers=12))

	conv_pb2_grpc.add_ColorConvServiceServicer_to_server(ColorConvSvc(), server)

	server.add_insecure_port("[::]:50051")
	server.start()
	server.wait_for_termination()
	pass

main()
