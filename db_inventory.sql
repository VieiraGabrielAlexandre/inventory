--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2
-- Dumped by pg_dump version 14.2

-- Started on 2022-04-02 17:10:50

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 210 (class 1259 OID 16396)
-- Name: produtos; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.produtos (
    id integer NOT NULL,
    nome character varying,
    descricao character varying,
    preco numeric,
    quantidade integer
);


ALTER TABLE public.produtos OWNER TO postgres;

--
-- TOC entry 209 (class 1259 OID 16395)
-- Name: produtos_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.produtos_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.produtos_id_seq OWNER TO postgres;

--
-- TOC entry 3313 (class 0 OID 0)
-- Dependencies: 209
-- Name: produtos_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.produtos_id_seq OWNED BY public.produtos.id;


--
-- TOC entry 3164 (class 2604 OID 16399)
-- Name: produtos id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.produtos ALTER COLUMN id SET DEFAULT nextval('public.produtos_id_seq'::regclass);


--
-- TOC entry 3307 (class 0 OID 16396)
-- Dependencies: 210
-- Data for Name: produtos; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.produtos (id, nome, descricao, preco, quantidade) VALUES (1, 'Macbook PRO 2012', 'Macbook usado com 8gb', 4000.00, 1);
INSERT INTO public.produtos (id, nome, descricao, preco, quantidade) VALUES (2, 'Macbook Air 2021', 'Macbook novo nunca usado', 8000, 5);
INSERT INTO public.produtos (id, nome, descricao, preco, quantidade) VALUES (3, 'Iphone 13', 'Iphone 13 novo, sem detalhes', 9000, 10);
INSERT INTO public.produtos (id, nome, descricao, preco, quantidade) VALUES (4, 'Magic Mouse', 'Teclado Apple raro', 15000, 1);


--
-- TOC entry 3314 (class 0 OID 0)
-- Dependencies: 209
-- Name: produtos_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.produtos_id_seq', 4, true);


--
-- TOC entry 3166 (class 2606 OID 16403)
-- Name: produtos produtos_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.produtos
    ADD CONSTRAINT produtos_pkey PRIMARY KEY (id);


-- Completed on 2022-04-02 17:10:50

--
-- PostgreSQL database dump complete
--

