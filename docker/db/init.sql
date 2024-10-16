-- Create Database
CREATE DATABASE recruitment;

\c recruitment;

-- Enum
CREATE TYPE role AS ENUM ('applicant', 'admin');