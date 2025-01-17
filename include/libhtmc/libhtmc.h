// MIT License
//
// Copyright (c) 2024 Alessandro Salerno
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

#pragma once
#include <stdarg.h>
#include <stdbool.h>
#include <stdio.h>

typedef enum htmc_handover_variant {
  HTMC_BASE_HANDOVER
} htmc_handover_variant_t;

typedef struct htmc_handover htmc_handover_t;
typedef struct htmc_handover {
  const htmc_handover_variant_t variant_id;
  int (*vprintf)(htmc_handover_t *handover, const char *fmt, va_list args);
  int (*puts)(htmc_handover_t *handover, const char *s);
  int (*query_vscanf)(htmc_handover_t *handover, const char *fmt, va_list args);
  int (*form_vscanf)(htmc_handover_t *handover, const char *fmt, va_list args);
  void *(*alloc)(htmc_handover_t *handover, size_t nbytes);
  void (*free)(htmc_handover_t *handover, void *ptr);

  const char *request_method;
  const char *query_string;
  size_t      content_length;
  const char *content_type;
  const char *request_body;
} htmc_handover_t;

void  htmc_bind(htmc_handover_t *handover);
int   htmc_printf(const char *fmt, ...);
int   htmc_vprintf(const char *fmt, va_list args);
int   htmc_puts(const char *s);
int   htmc_query_scanf(const char *fmt, ...);
int   htmc_query_vscanf(const char *fmt, va_list args);
int   htmc_form_scanf(const char *fmt, ...);
int   htmc_form_vscafn(const char *fmt, va_list args);
void *htmc_alloc(size_t nbytes);
void  htmc_free(void *ptr);
void  htmc_error(const char *fmt, ...);
void  htmc_verror(const char *fmt, va_list args);
