#!/bin/bash
sudo -i -u postgres psql < setup.sql # Это для Linux
psql -d postgres < setup.sql # Это для MacOs