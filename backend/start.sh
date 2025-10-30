#!/bin/bash
# Forzar a Go a usar su resolutor interno para evitar el problema de IPv6 de Render
GODEBUG=netdns=go ./main