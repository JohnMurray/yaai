// Code generated from yaai/YaaiLexer.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 49, 319,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 7,
	21, 221, 10, 21, 12, 21, 14, 21, 224, 11, 21, 3, 22, 3, 22, 3, 22, 6, 22,
	229, 10, 22, 13, 22, 14, 22, 230, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 6,
	24, 238, 10, 24, 13, 24, 14, 24, 239, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26,
	3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 30, 3,
	30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33,
	3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3,
	38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 43,
	3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3,
	48, 3, 48, 3, 48, 3, 48, 7, 48, 303, 10, 48, 12, 48, 14, 48, 306, 11, 48,
	3, 48, 3, 48, 3, 49, 6, 49, 311, 10, 49, 13, 49, 14, 49, 312, 3, 49, 5,
	49, 316, 10, 49, 3, 49, 3, 49, 2, 2, 50, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7,
	13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31,
	17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 2, 45, 23, 47, 24, 49,
	25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67,
	34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85,
	43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 3, 2, 8, 4, 2, 67,
	92, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 36, 36, 3, 2,
	50, 59, 4, 2, 11, 11, 34, 34, 4, 2, 12, 12, 15, 15, 2, 324, 2, 3, 3, 2,
	2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2,
	2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3,
	2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27,
	3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2,
	35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2,
	2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2,
	2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2,
	2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3,
	2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75,
	3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2,
	83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2,
	2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2,
	2, 3, 99, 3, 2, 2, 2, 5, 105, 3, 2, 2, 2, 7, 109, 3, 2, 2, 2, 9, 114, 3,
	2, 2, 2, 11, 119, 3, 2, 2, 2, 13, 129, 3, 2, 2, 2, 15, 137, 3, 2, 2, 2,
	17, 145, 3, 2, 2, 2, 19, 150, 3, 2, 2, 2, 21, 157, 3, 2, 2, 2, 23, 162,
	3, 2, 2, 2, 25, 172, 3, 2, 2, 2, 27, 176, 3, 2, 2, 2, 29, 180, 3, 2, 2,
	2, 31, 186, 3, 2, 2, 2, 33, 192, 3, 2, 2, 2, 35, 197, 3, 2, 2, 2, 37, 204,
	3, 2, 2, 2, 39, 211, 3, 2, 2, 2, 41, 218, 3, 2, 2, 2, 43, 228, 3, 2, 2,
	2, 45, 232, 3, 2, 2, 2, 47, 237, 3, 2, 2, 2, 49, 241, 3, 2, 2, 2, 51, 244,
	3, 2, 2, 2, 53, 246, 3, 2, 2, 2, 55, 249, 3, 2, 2, 2, 57, 252, 3, 2, 2,
	2, 59, 255, 3, 2, 2, 2, 61, 258, 3, 2, 2, 2, 63, 261, 3, 2, 2, 2, 65, 264,
	3, 2, 2, 2, 67, 267, 3, 2, 2, 2, 69, 270, 3, 2, 2, 2, 71, 272, 3, 2, 2,
	2, 73, 274, 3, 2, 2, 2, 75, 276, 3, 2, 2, 2, 77, 278, 3, 2, 2, 2, 79, 280,
	3, 2, 2, 2, 81, 282, 3, 2, 2, 2, 83, 284, 3, 2, 2, 2, 85, 286, 3, 2, 2,
	2, 87, 288, 3, 2, 2, 2, 89, 290, 3, 2, 2, 2, 91, 292, 3, 2, 2, 2, 93, 294,
	3, 2, 2, 2, 95, 298, 3, 2, 2, 2, 97, 315, 3, 2, 2, 2, 99, 100, 7, 99, 2,
	2, 100, 101, 7, 101, 2, 2, 101, 102, 7, 118, 2, 2, 102, 103, 7, 113, 2,
	2, 103, 104, 7, 116, 2, 2, 104, 4, 3, 2, 2, 2, 105, 106, 7, 104, 2, 2,
	106, 107, 7, 113, 2, 2, 107, 108, 7, 116, 2, 2, 108, 6, 3, 2, 2, 2, 109,
	110, 7, 104, 2, 2, 110, 111, 7, 119, 2, 2, 111, 112, 7, 112, 2, 2, 112,
	113, 7, 101, 2, 2, 113, 8, 3, 2, 2, 2, 114, 115, 7, 107, 2, 2, 115, 116,
	7, 112, 2, 2, 116, 117, 7, 107, 2, 2, 117, 118, 7, 118, 2, 2, 118, 10,
	3, 2, 2, 2, 119, 120, 7, 107, 2, 2, 120, 121, 7, 112, 2, 2, 121, 122, 7,
	118, 2, 2, 122, 123, 7, 103, 2, 2, 123, 124, 7, 116, 2, 2, 124, 125, 7,
	104, 2, 2, 125, 126, 7, 99, 2, 2, 126, 127, 7, 101, 2, 2, 127, 128, 7,
	103, 2, 2, 128, 12, 3, 2, 2, 2, 129, 130, 7, 114, 2, 2, 130, 131, 7, 99,
	2, 2, 131, 132, 7, 101, 2, 2, 132, 133, 7, 109, 2, 2, 133, 134, 7, 99,
	2, 2, 134, 135, 7, 105, 2, 2, 135, 136, 7, 103, 2, 2, 136, 14, 3, 2, 2,
	2, 137, 138, 7, 116, 2, 2, 138, 139, 7, 103, 2, 2, 139, 140, 7, 101, 2,
	2, 140, 141, 7, 103, 2, 2, 141, 142, 7, 107, 2, 2, 142, 143, 7, 120, 2,
	2, 143, 144, 7, 103, 2, 2, 144, 16, 3, 2, 2, 2, 145, 146, 7, 117, 2, 2,
	146, 147, 7, 103, 2, 2, 147, 148, 7, 110, 2, 2, 148, 149, 7, 104, 2, 2,
	149, 18, 3, 2, 2, 2, 150, 151, 7, 117, 2, 2, 151, 152, 7, 118, 2, 2, 152,
	153, 7, 116, 2, 2, 153, 154, 7, 119, 2, 2, 154, 155, 7, 101, 2, 2, 155,
	156, 7, 118, 2, 2, 156, 20, 3, 2, 2, 2, 157, 158, 7, 118, 2, 2, 158, 159,
	7, 123, 2, 2, 159, 160, 7, 114, 2, 2, 160, 161, 7, 103, 2, 2, 161, 22,
	3, 2, 2, 2, 162, 163, 7, 119, 2, 2, 163, 164, 7, 112, 2, 2, 164, 165, 7,
	106, 2, 2, 165, 166, 7, 99, 2, 2, 166, 167, 7, 112, 2, 2, 167, 168, 7,
	102, 2, 2, 168, 169, 7, 110, 2, 2, 169, 170, 7, 103, 2, 2, 170, 171, 7,
	102, 2, 2, 171, 24, 3, 2, 2, 2, 172, 173, 7, 120, 2, 2, 173, 174, 7, 99,
	2, 2, 174, 175, 7, 116, 2, 2, 175, 26, 3, 2, 2, 2, 176, 177, 7, 107, 2,
	2, 177, 178, 7, 112, 2, 2, 178, 179, 7, 118, 2, 2, 179, 28, 3, 2, 2, 2,
	180, 181, 7, 107, 2, 2, 181, 182, 7, 112, 2, 2, 182, 183, 7, 118, 2, 2,
	183, 184, 7, 53, 2, 2, 184, 185, 7, 52, 2, 2, 185, 30, 3, 2, 2, 2, 186,
	187, 7, 107, 2, 2, 187, 188, 7, 112, 2, 2, 188, 189, 7, 118, 2, 2, 189,
	190, 7, 56, 2, 2, 190, 191, 7, 54, 2, 2, 191, 32, 3, 2, 2, 2, 192, 193,
	7, 119, 2, 2, 193, 194, 7, 107, 2, 2, 194, 195, 7, 112, 2, 2, 195, 196,
	7, 118, 2, 2, 196, 34, 3, 2, 2, 2, 197, 198, 7, 119, 2, 2, 198, 199, 7,
	107, 2, 2, 199, 200, 7, 112, 2, 2, 200, 201, 7, 118, 2, 2, 201, 202, 7,
	53, 2, 2, 202, 203, 7, 52, 2, 2, 203, 36, 3, 2, 2, 2, 204, 205, 7, 119,
	2, 2, 205, 206, 7, 107, 2, 2, 206, 207, 7, 112, 2, 2, 207, 208, 7, 118,
	2, 2, 208, 209, 7, 56, 2, 2, 209, 210, 7, 54, 2, 2, 210, 38, 3, 2, 2, 2,
	211, 212, 7, 117, 2, 2, 212, 213, 7, 118, 2, 2, 213, 214, 7, 116, 2, 2,
	214, 215, 7, 107, 2, 2, 215, 216, 7, 112, 2, 2, 216, 217, 7, 105, 2, 2,
	217, 40, 3, 2, 2, 2, 218, 222, 9, 2, 2, 2, 219, 221, 9, 3, 2, 2, 220, 219,
	3, 2, 2, 2, 221, 224, 3, 2, 2, 2, 222, 220, 3, 2, 2, 2, 222, 223, 3, 2,
	2, 2, 223, 42, 3, 2, 2, 2, 224, 222, 3, 2, 2, 2, 225, 226, 7, 94, 2, 2,
	226, 229, 7, 36, 2, 2, 227, 229, 10, 4, 2, 2, 228, 225, 3, 2, 2, 2, 228,
	227, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 228, 3, 2, 2, 2, 230, 231,
	3, 2, 2, 2, 231, 44, 3, 2, 2, 2, 232, 233, 7, 36, 2, 2, 233, 234, 5, 43,
	22, 2, 234, 235, 7, 36, 2, 2, 235, 46, 3, 2, 2, 2, 236, 238, 9, 5, 2, 2,
	237, 236, 3, 2, 2, 2, 238, 239, 3, 2, 2, 2, 239, 237, 3, 2, 2, 2, 239,
	240, 3, 2, 2, 2, 240, 48, 3, 2, 2, 2, 241, 242, 7, 60, 2, 2, 242, 243,
	7, 63, 2, 2, 243, 50, 3, 2, 2, 2, 244, 245, 7, 63, 2, 2, 245, 52, 3, 2,
	2, 2, 246, 247, 7, 45, 2, 2, 247, 248, 7, 45, 2, 2, 248, 54, 3, 2, 2, 2,
	249, 250, 7, 47, 2, 2, 250, 251, 7, 47, 2, 2, 251, 56, 3, 2, 2, 2, 252,
	253, 7, 45, 2, 2, 253, 254, 7, 63, 2, 2, 254, 58, 3, 2, 2, 2, 255, 256,
	7, 47, 2, 2, 256, 257, 7, 63, 2, 2, 257, 60, 3, 2, 2, 2, 258, 259, 7, 63,
	2, 2, 259, 260, 7, 63, 2, 2, 260, 62, 3, 2, 2, 2, 261, 262, 7, 35, 2, 2,
	262, 263, 7, 63, 2, 2, 263, 64, 3, 2, 2, 2, 264, 265, 7, 62, 2, 2, 265,
	266, 7, 63, 2, 2, 266, 66, 3, 2, 2, 2, 267, 268, 7, 64, 2, 2, 268, 269,
	7, 63, 2, 2, 269, 68, 3, 2, 2, 2, 270, 271, 7, 62, 2, 2, 271, 70, 3, 2,
	2, 2, 272, 273, 7, 64, 2, 2, 273, 72, 3, 2, 2, 2, 274, 275, 7, 125, 2,
	2, 275, 74, 3, 2, 2, 2, 276, 277, 7, 127, 2, 2, 277, 76, 3, 2, 2, 2, 278,
	279, 7, 42, 2, 2, 279, 78, 3, 2, 2, 2, 280, 281, 7, 43, 2, 2, 281, 80,
	3, 2, 2, 2, 282, 283, 7, 48, 2, 2, 283, 82, 3, 2, 2, 2, 284, 285, 7, 45,
	2, 2, 285, 84, 3, 2, 2, 2, 286, 287, 7, 47, 2, 2, 287, 86, 3, 2, 2, 2,
	288, 289, 7, 44, 2, 2, 289, 88, 3, 2, 2, 2, 290, 291, 7, 49, 2, 2, 291,
	90, 3, 2, 2, 2, 292, 293, 7, 46, 2, 2, 293, 92, 3, 2, 2, 2, 294, 295, 9,
	6, 2, 2, 295, 296, 3, 2, 2, 2, 296, 297, 8, 47, 2, 2, 297, 94, 3, 2, 2,
	2, 298, 299, 7, 49, 2, 2, 299, 300, 7, 49, 2, 2, 300, 304, 3, 2, 2, 2,
	301, 303, 10, 7, 2, 2, 302, 301, 3, 2, 2, 2, 303, 306, 3, 2, 2, 2, 304,
	302, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 305, 307, 3, 2, 2, 2, 306, 304,
	3, 2, 2, 2, 307, 308, 8, 48, 3, 2, 308, 96, 3, 2, 2, 2, 309, 311, 9, 7,
	2, 2, 310, 309, 3, 2, 2, 2, 311, 312, 3, 2, 2, 2, 312, 310, 3, 2, 2, 2,
	312, 313, 3, 2, 2, 2, 313, 316, 3, 2, 2, 2, 314, 316, 7, 61, 2, 2, 315,
	310, 3, 2, 2, 2, 315, 314, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 318,
	8, 49, 4, 2, 318, 98, 3, 2, 2, 2, 10, 2, 222, 228, 230, 239, 304, 312,
	315, 5, 8, 2, 2, 2, 3, 2, 4, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'actor'", "'for'", "'func'", "'init'", "'interface'", "'package'",
	"'receive'", "'self'", "'struct'", "'type'", "'unhandled'", "'var'", "'int'",
	"'int32'", "'int64'", "'uint'", "'uint32'", "'uint64'", "'string'", "",
	"", "", "':='", "'='", "'++'", "'--'", "'+='", "'-='", "'=='", "'!='",
	"'<='", "'>='", "'<'", "'>'", "'{'", "'}'", "'('", "')'", "'.'", "'+'",
	"'-'", "'*'", "'/'", "','",
}

var lexerSymbolicNames = []string{
	"", "ACTOR", "FOR", "FUNC", "INIT", "INTERFACE", "PACKAGE", "RECEIVE",
	"SELF", "STRUCT", "TYPE", "UNHANDLED", "VAR", "T_INT", "T_INT32", "T_INT64",
	"T_UINT", "T_UINT32", "T_UINT64", "T_STRING", "IDENTIFIER", "STRING_LITERAL",
	"NUMERIC_LITERAL", "VAR_INITIALIZER", "ASSIGNMENT", "INCR", "DECR", "PLUS_EQUAL",
	"MINUS_EQUAL", "EQUAL_EQUAL", "NOT_EQUAL", "LESS_THAN_EQUAL", "GREATER_THAN_EQUAL",
	"LESS_THAN", "GREATER_THAN", "L_BRACKET", "R_BRACKET", "L_PAREN", "R_PAREN",
	"DOT", "PLUS", "MINUS", "STAR", "F_SLASH", "COMMA", "NB_WS", "LINE_COMMENT",
	"EOS",
}

var lexerRuleNames = []string{
	"ACTOR", "FOR", "FUNC", "INIT", "INTERFACE", "PACKAGE", "RECEIVE", "SELF",
	"STRUCT", "TYPE", "UNHANDLED", "VAR", "T_INT", "T_INT32", "T_INT64", "T_UINT",
	"T_UINT32", "T_UINT64", "T_STRING", "IDENTIFIER", "STRING_BODY", "STRING_LITERAL",
	"NUMERIC_LITERAL", "VAR_INITIALIZER", "ASSIGNMENT", "INCR", "DECR", "PLUS_EQUAL",
	"MINUS_EQUAL", "EQUAL_EQUAL", "NOT_EQUAL", "LESS_THAN_EQUAL", "GREATER_THAN_EQUAL",
	"LESS_THAN", "GREATER_THAN", "L_BRACKET", "R_BRACKET", "L_PAREN", "R_PAREN",
	"DOT", "PLUS", "MINUS", "STAR", "F_SLASH", "COMMA", "NB_WS", "LINE_COMMENT",
	"EOS",
}

type YaaiLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewYaaiLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *YaaiLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewYaaiLexer(input antlr.CharStream) *YaaiLexer {
	l := new(YaaiLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "YaaiLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// YaaiLexer tokens.
const (
	YaaiLexerACTOR              = 1
	YaaiLexerFOR                = 2
	YaaiLexerFUNC               = 3
	YaaiLexerINIT               = 4
	YaaiLexerINTERFACE          = 5
	YaaiLexerPACKAGE            = 6
	YaaiLexerRECEIVE            = 7
	YaaiLexerSELF               = 8
	YaaiLexerSTRUCT             = 9
	YaaiLexerTYPE               = 10
	YaaiLexerUNHANDLED          = 11
	YaaiLexerVAR                = 12
	YaaiLexerT_INT              = 13
	YaaiLexerT_INT32            = 14
	YaaiLexerT_INT64            = 15
	YaaiLexerT_UINT             = 16
	YaaiLexerT_UINT32           = 17
	YaaiLexerT_UINT64           = 18
	YaaiLexerT_STRING           = 19
	YaaiLexerIDENTIFIER         = 20
	YaaiLexerSTRING_LITERAL     = 21
	YaaiLexerNUMERIC_LITERAL    = 22
	YaaiLexerVAR_INITIALIZER    = 23
	YaaiLexerASSIGNMENT         = 24
	YaaiLexerINCR               = 25
	YaaiLexerDECR               = 26
	YaaiLexerPLUS_EQUAL         = 27
	YaaiLexerMINUS_EQUAL        = 28
	YaaiLexerEQUAL_EQUAL        = 29
	YaaiLexerNOT_EQUAL          = 30
	YaaiLexerLESS_THAN_EQUAL    = 31
	YaaiLexerGREATER_THAN_EQUAL = 32
	YaaiLexerLESS_THAN          = 33
	YaaiLexerGREATER_THAN       = 34
	YaaiLexerL_BRACKET          = 35
	YaaiLexerR_BRACKET          = 36
	YaaiLexerL_PAREN            = 37
	YaaiLexerR_PAREN            = 38
	YaaiLexerDOT                = 39
	YaaiLexerPLUS               = 40
	YaaiLexerMINUS              = 41
	YaaiLexerSTAR               = 42
	YaaiLexerF_SLASH            = 43
	YaaiLexerCOMMA              = 44
	YaaiLexerNB_WS              = 45
	YaaiLexerLINE_COMMENT       = 46
	YaaiLexerEOS                = 47
)