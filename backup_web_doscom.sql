--
-- PostgreSQL database dump
--

\restrict JHgEa2pXsuyJWQOQc7vstfhyb1phOEWgOUpTett2M0cNLpegp6s7B23RNHuIdrz

-- Dumped from database version 16.14
-- Dumped by pg_dump version 16.14

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

--
-- Name: blog_category; Type: TYPE; Schema: public; Owner: iko
--

CREATE TYPE public.blog_category AS ENUM (
    'event',
    'seminar',
    'collaboration',
    'education',
    'technology',
    'work',
    'activity',
    'sharing'
);


ALTER TYPE public.blog_category OWNER TO iko;

--
-- Name: blog_status; Type: TYPE; Schema: public; Owner: iko
--

CREATE TYPE public.blog_status AS ENUM (
    'draft',
    'published',
    'scheduled',
    'unpublished',
    'rejected',
    'pending_review'
);


ALTER TYPE public.blog_status OWNER TO iko;

--
-- Name: work_status; Type: TYPE; Schema: public; Owner: iko
--

CREATE TYPE public.work_status AS ENUM (
    'draft',
    'published',
    'unpublished',
    'rejected',
    'pending_review'
);


ALTER TYPE public.work_status OWNER TO iko;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: blog; Type: TABLE; Schema: public; Owner: iko
--

CREATE TABLE public.blog (
    id integer NOT NULL,
    author_id integer,
    title character varying(255) NOT NULL,
    slug character varying(255) NOT NULL,
    content text,
    thumbnail_url text,
    kategori public.blog_category[] NOT NULL,
    published_at timestamp without time zone,
    status public.blog_status DEFAULT 'draft'::public.blog_status NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.blog OWNER TO iko;

--
-- Name: blog_gallery; Type: TABLE; Schema: public; Owner: iko
--

CREATE TABLE public.blog_gallery (
    id integer NOT NULL,
    id_blog integer,
    id_gallery integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.blog_gallery OWNER TO iko;

--
-- Name: blog_gallery_id_seq; Type: SEQUENCE; Schema: public; Owner: iko
--

CREATE SEQUENCE public.blog_gallery_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.blog_gallery_id_seq OWNER TO iko;

--
-- Name: blog_gallery_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: iko
--

ALTER SEQUENCE public.blog_gallery_id_seq OWNED BY public.blog_gallery.id;


--
-- Name: blog_id_seq; Type: SEQUENCE; Schema: public; Owner: iko
--

CREATE SEQUENCE public.blog_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.blog_id_seq OWNER TO iko;

--
-- Name: blog_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: iko
--

ALTER SEQUENCE public.blog_id_seq OWNED BY public.blog.id;


--
-- Name: file_uploads; Type: TABLE; Schema: public; Owner: iko
--

CREATE TABLE public.file_uploads (
    id integer NOT NULL,
    user_id integer NOT NULL,
    category character varying(50) NOT NULL,
    original_filename character varying(255) NOT NULL,
    stored_filename character varying(255) NOT NULL,
    file_size bigint NOT NULL,
    content_type character varying(100) NOT NULL,
    file_url text NOT NULL,
    uploaded_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT file_uploads_category_check CHECK (((category)::text = ANY ((ARRAY['gallery'::character varying, 'blog'::character varying, 'work'::character varying, 'pengurus'::character varying])::text[])))
);


ALTER TABLE public.file_uploads OWNER TO iko;

--
-- Name: file_uploads_id_seq; Type: SEQUENCE; Schema: public; Owner: iko
--

CREATE SEQUENCE public.file_uploads_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.file_uploads_id_seq OWNER TO iko;

--
-- Name: file_uploads_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: iko
--

ALTER SEQUENCE public.file_uploads_id_seq OWNED BY public.file_uploads.id;


--
-- Name: gallery; Type: TABLE; Schema: public; Owner: iko
--

CREATE TABLE public.gallery (
    id integer NOT NULL,
    id_users integer,
    file_upload_id integer,
    gallery_name character varying(150) NOT NULL,
    gallery_type character varying(50),
    description text,
    event_date timestamp without time zone,
    file_url text NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT gallery_gallery_type_check CHECK (((gallery_type)::text = ANY ((ARRAY['fun'::character varying, 'proker'::character varying, 'achievment'::character varying, 'work'::character varying, 'activity'::character varying, 'blog'::character varying, 'pengurus'::character varying, 'etc'::character varying])::text[])))
);


ALTER TABLE public.gallery OWNER TO iko;

--
-- Name: gallery_id_seq; Type: SEQUENCE; Schema: public; Owner: iko
--

CREATE SEQUENCE public.gallery_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.gallery_id_seq OWNER TO iko;

--
-- Name: gallery_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: iko
--

ALTER SEQUENCE public.gallery_id_seq OWNED BY public.gallery.id;


--
-- Name: history_photos; Type: TABLE; Schema: public; Owner: iko
--

CREATE TABLE public.history_photos (
    id integer NOT NULL,
    id_history integer,
    imager_url text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.history_photos OWNER TO iko;

--
-- Name: history_photos_id_seq; Type: SEQUENCE; Schema: public; Owner: iko
--

CREATE SEQUENCE public.history_photos_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.history_photos_id_seq OWNER TO iko;

--
-- Name: history_photos_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: iko
--

ALTER SEQUENCE public.history_photos_id_seq OWNED BY public.history_photos.id;


--
-- Name: history_timeline; Type: TABLE; Schema: public; Owner: iko
--

CREATE TABLE public.history_timeline (
    id integer NOT NULL,
    id_author integer,
    title character varying(255) NOT NULL,
    year character varying(50) NOT NULL,
    description text,
    display_order integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.history_timeline OWNER TO iko;

--
-- Name: history_timeline_id_seq; Type: SEQUENCE; Schema: public; Owner: iko
--

CREATE SEQUENCE public.history_timeline_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.history_timeline_id_seq OWNER TO iko;

--
-- Name: history_timeline_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: iko
--

ALTER SEQUENCE public.history_timeline_id_seq OWNED BY public.history_timeline.id;


--
-- Name: pengurus; Type: TABLE; Schema: public; Owner: iko
--

CREATE TABLE public.pengurus (
    id integer NOT NULL,
    id_user integer,
    photo_url text,
    email character varying(100) NOT NULL,
    divisi character varying(50),
    name character varying(150) NOT NULL,
    "position" character varying(100),
    start_periode_year integer NOT NULL,
    end_periode_year integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pengurus_divisi_check CHECK (((divisi)::text = ANY ((ARRAY['bph'::character varying, 'pemro'::character varying, 'jaringan'::character varying, 'medcrev'::character varying, 'data'::character varying])::text[]))),
    CONSTRAINT pengurus_position_check CHECK ((("position")::text = ANY ((ARRAY['ketum'::character varying, 'sdm'::character varying, 'pr'::character varying, 'pm'::character varying, 'pmAng'::character varying, 'sekum'::character varying, 'bendum'::character varying, 'sekAng'::character varying, 'bendAng'::character varying, 'KoorPemro'::character varying, 'KoorJaringan'::character varying, 'KoorMedcrev'::character varying, 'KoorData'::character varying, 'anggotaAktif'::character varying, 'PemroAng'::character varying, 'JaringanAng'::character varying, 'MedcrevAng'::character varying, 'DataAng'::character varying])::text[])))
);


ALTER TABLE public.pengurus OWNER TO iko;

--
-- Name: pengurus_id_seq; Type: SEQUENCE; Schema: public; Owner: iko
--

CREATE SEQUENCE public.pengurus_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.pengurus_id_seq OWNER TO iko;

--
-- Name: pengurus_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: iko
--

ALTER SEQUENCE public.pengurus_id_seq OWNED BY public.pengurus.id;


--
-- Name: pengurus_sosmed; Type: TABLE; Schema: public; Owner: iko
--

CREATE TABLE public.pengurus_sosmed (
    id integer NOT NULL,
    pengurus_id integer NOT NULL,
    platform character varying(50) NOT NULL,
    username character varying(100),
    url text,
    is_primary boolean DEFAULT false,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.pengurus_sosmed OWNER TO iko;

--
-- Name: pengurus_sosmed_id_seq; Type: SEQUENCE; Schema: public; Owner: iko
--

CREATE SEQUENCE public.pengurus_sosmed_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.pengurus_sosmed_id_seq OWNER TO iko;

--
-- Name: pengurus_sosmed_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: iko
--

ALTER SEQUENCE public.pengurus_sosmed_id_seq OWNED BY public.pengurus_sosmed.id;


--
-- Name: refresh_token; Type: TABLE; Schema: public; Owner: iko
--

CREATE TABLE public.refresh_token (
    id integer NOT NULL,
    user_id integer NOT NULL,
    token character varying(255) NOT NULL,
    expires timestamp without time zone NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.refresh_token OWNER TO iko;

--
-- Name: refresh_token_id_seq; Type: SEQUENCE; Schema: public; Owner: iko
--

CREATE SEQUENCE public.refresh_token_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.refresh_token_id_seq OWNER TO iko;

--
-- Name: refresh_token_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: iko
--

ALTER SEQUENCE public.refresh_token_id_seq OWNED BY public.refresh_token.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: iko
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO iko;

--
-- Name: users; Type: TABLE; Schema: public; Owner: iko
--

CREATE TABLE public.users (
    id integer NOT NULL,
    email character varying(100),
    username character varying(100) NOT NULL,
    role character varying(50),
    password character varying(100) NOT NULL,
    full_name character varying(100),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT users_role_check CHECK (((role)::text = ANY ((ARRAY['SuperAdmin'::character varying, 'KoorPemro'::character varying, 'KoorJaringan'::character varying, 'KoorData'::character varying, 'KoorMedcrev'::character varying, 'BPH'::character varying, 'pemroAnggota'::character varying, 'jaringanAnggota'::character varying, 'medcrevAnggota'::character varying, 'dataAnggota'::character varying, 'BPHAnggota'::character varying])::text[])))
);


ALTER TABLE public.users OWNER TO iko;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: iko
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO iko;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: iko
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: work; Type: TABLE; Schema: public; Owner: iko
--

CREATE TABLE public.work (
    id integer NOT NULL,
    pengurus_id integer,
    title character varying(300) NOT NULL,
    tagline character varying(100),
    description text,
    slug text,
    project_type character varying(50),
    technologies text[],
    project_date timestamp without time zone,
    image_url text,
    status public.work_status DEFAULT 'draft'::public.work_status NOT NULL,
    division character varying(50),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.work OWNER TO iko;

--
-- Name: work_gallery; Type: TABLE; Schema: public; Owner: iko
--

CREATE TABLE public.work_gallery (
    id integer NOT NULL,
    id_work integer,
    id_gallery integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.work_gallery OWNER TO iko;

--
-- Name: work_gallery_id_seq; Type: SEQUENCE; Schema: public; Owner: iko
--

CREATE SEQUENCE public.work_gallery_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.work_gallery_id_seq OWNER TO iko;

--
-- Name: work_gallery_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: iko
--

ALTER SEQUENCE public.work_gallery_id_seq OWNED BY public.work_gallery.id;


--
-- Name: work_id_seq; Type: SEQUENCE; Schema: public; Owner: iko
--

CREATE SEQUENCE public.work_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.work_id_seq OWNER TO iko;

--
-- Name: work_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: iko
--

ALTER SEQUENCE public.work_id_seq OWNED BY public.work.id;


--
-- Name: blog id; Type: DEFAULT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.blog ALTER COLUMN id SET DEFAULT nextval('public.blog_id_seq'::regclass);


--
-- Name: blog_gallery id; Type: DEFAULT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.blog_gallery ALTER COLUMN id SET DEFAULT nextval('public.blog_gallery_id_seq'::regclass);


--
-- Name: file_uploads id; Type: DEFAULT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.file_uploads ALTER COLUMN id SET DEFAULT nextval('public.file_uploads_id_seq'::regclass);


--
-- Name: gallery id; Type: DEFAULT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.gallery ALTER COLUMN id SET DEFAULT nextval('public.gallery_id_seq'::regclass);


--
-- Name: history_photos id; Type: DEFAULT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.history_photos ALTER COLUMN id SET DEFAULT nextval('public.history_photos_id_seq'::regclass);


--
-- Name: history_timeline id; Type: DEFAULT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.history_timeline ALTER COLUMN id SET DEFAULT nextval('public.history_timeline_id_seq'::regclass);


--
-- Name: pengurus id; Type: DEFAULT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.pengurus ALTER COLUMN id SET DEFAULT nextval('public.pengurus_id_seq'::regclass);


--
-- Name: pengurus_sosmed id; Type: DEFAULT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.pengurus_sosmed ALTER COLUMN id SET DEFAULT nextval('public.pengurus_sosmed_id_seq'::regclass);


--
-- Name: refresh_token id; Type: DEFAULT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.refresh_token ALTER COLUMN id SET DEFAULT nextval('public.refresh_token_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: work id; Type: DEFAULT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.work ALTER COLUMN id SET DEFAULT nextval('public.work_id_seq'::regclass);


--
-- Name: work_gallery id; Type: DEFAULT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.work_gallery ALTER COLUMN id SET DEFAULT nextval('public.work_gallery_id_seq'::regclass);


--
-- Data for Name: blog; Type: TABLE DATA; Schema: public; Owner: iko
--

COPY public.blog (id, author_id, title, slug, content, thumbnail_url, kategori, published_at, status, created_at, updated_at) FROM stdin;
1	3	Pengenalan Open Source	pengenalan-open-source	Open Source adalah masa depan perangkat lunak.		{technology,education}	2026-05-27 06:26:32.861531	published	2026-05-27 06:26:32.861553	2026-05-27 06:26:32.861553
2	3	Cara Membuat Backend Go	cara-membuat-backend-go	Tutorial membuat REST API dengan Gin dan GORM.		{technology}	2026-05-27 06:26:32.861531	published	2026-05-27 06:26:32.866204	2026-05-27 06:26:32.866204
3	3	Mengenal JWT Authentication di Golang	mengenal-jwt-authentication-di-golang	Panduan implementasi JWT authentication menggunakan Gin framework.		{technology,education}	2026-05-27 06:26:32.861531	published	2026-05-27 06:26:32.867314	2026-05-27 06:26:32.867314
4	3	Belajar PostgreSQL untuk Backend Developer	belajar-postgresql-untuk-backend-developer	Dasar penggunaan PostgreSQL untuk aplikasi backend modern.		{education,technology}	2026-05-27 06:26:32.861531	published	2026-05-27 06:26:32.868174	2026-05-27 06:26:32.868174
5	3	Deploy Aplikasi Go ke VPS Linux	deploy-aplikasi-go-ke-vps-linux	Tutorial deploy aplikasi Golang menggunakan svstemd dan Nginx.		{technology,work}	\N	draft	2026-05-27 06:26:32.868866	2026-05-27 06:26:32.868866
6	3	Membuat CRUD API dengan Gin dan GORM	membuat-crud-api-dengan-gin-dan-gorm	Step by step membuat CRUD API menggunakan Golang, Gin, dan GORM.		{activity,technology}	2026-05-27 06:26:32.861531	published	2026-05-27 06:26:32.869786	2026-05-27 06:26:32.869787
7	3	Pengenalan Docker untuk Developer	pengenalan-docker-untuk-developer	Belajar dasar Docker dan containerization untuk development workflow.		{education,technology,work}	2026-05-27 06:26:32.861531	published	2026-05-27 06:26:32.870918	2026-05-27 06:26:32.870919
\.


--
-- Data for Name: blog_gallery; Type: TABLE DATA; Schema: public; Owner: iko
--

COPY public.blog_gallery (id, id_blog, id_gallery, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: file_uploads; Type: TABLE DATA; Schema: public; Owner: iko
--

COPY public.file_uploads (id, user_id, category, original_filename, stored_filename, file_size, content_type, file_url, uploaded_at, updated_at) FROM stdin;
1	7	pengurus	26-Photoroom.png	pengurus/20260528-279aec68-d4d0-4dc0-87a2-b78e69c1654d.png	188890	image/png	http://localhost:9000/doscom-uploads/pengurus/20260528-279aec68-d4d0-4dc0-87a2-b78e69c1654d.png	2026-05-28 19:00:40.052666	2026-05-28 19:00:40.052666
2	7	pengurus	26-Photoroom.png	pengurus/20260528-a9cf6483-1d74-404c-8daf-62da181289e7.png	188890	image/png	http://localhost:9000/doscom-uploads/pengurus/20260528-a9cf6483-1d74-404c-8daf-62da181289e7.png	2026-05-28 19:27:27.144906	2026-05-28 19:27:27.144906
3	7	pengurus	26-Photoroom.png	pengurus/20260528-a62f11d9-ca8f-4d0a-9c1c-2d9b78ecb197.png	188890	image/png	http://localhost:9000/doscom-uploads/pengurus/20260528-a62f11d9-ca8f-4d0a-9c1c-2d9b78ecb197.png	2026-05-28 19:31:04.249879	2026-05-28 19:31:04.249879
4	7	pengurus	26-Photoroom.png	pengurus/20260528-3f00b57f-4e2e-48e4-befc-9b509c2efbd0.png	188890	image/png	http://localhost:9000/doscom-uploads/pengurus/20260528-3f00b57f-4e2e-48e4-befc-9b509c2efbd0.png	2026-05-28 19:32:44.630629	2026-05-28 19:32:44.630629
5	7	pengurus	26-Photoroom.png	pengurus/20260529-5506a862-30cd-4749-8d82-13fc18e48bab.png	188890	image/png	http://localhost:9000/doscom-uploads/pengurus/20260529-5506a862-30cd-4749-8d82-13fc18e48bab.png	2026-05-29 11:46:52.571737	2026-05-29 11:46:52.571737
6	21	pengurus	26-Photoroom.png	pengurus/20260608-72cc7fd8-647e-4601-80e9-136b9444666b.png	188890	image/png	http://localhost:9000/doscom-uploads/pengurus/20260608-72cc7fd8-647e-4601-80e9-136b9444666b.png	2026-06-08 11:22:09.349068	2026-06-08 11:22:09.349068
7	21	pengurus	26-Photoroom.png	pengurus/20260608-ff1bf128-db0b-4b9a-a2b9-dd5c9965d23b.png	188890	image/png	http://localhost:9000/doscom-uploads/pengurus/20260608-ff1bf128-db0b-4b9a-a2b9-dd5c9965d23b.png	2026-06-08 11:23:05.262424	2026-06-08 11:23:05.262424
8	21	pengurus	26-Photoroom.png	pengurus/20260608-90d29552-a32e-4824-b3a2-672f759e78f9.png	188890	image/png	http://localhost:9000/doscom-uploads/pengurus/20260608-90d29552-a32e-4824-b3a2-672f759e78f9.png	2026-06-08 11:31:09.928386	2026-06-08 11:31:09.928386
9	21	pengurus	26-Photoroom.png	pengurus/20260608-244b2081-82f7-4150-9c98-2b77bc01678c.png	188890	image/png	http://localhost:9000/doscom-uploads/pengurus/20260608-244b2081-82f7-4150-9c98-2b77bc01678c.png	2026-06-08 12:23:11.033625	2026-06-08 12:23:11.033625
10	21	pengurus	26-Photoroom.png	pengurus/20260608-cef312e1-0299-4619-9d05-e6721b7ae9c6.png	188890	image/png	http://localhost:9000/doscom-uploads/pengurus/20260608-cef312e1-0299-4619-9d05-e6721b7ae9c6.png	2026-06-08 12:24:31.895831	2026-06-08 12:24:31.895831
11	21	pengurus	26-Photoroom.png	pengurus/20260608-cb913c09-eef7-4f19-a7fd-ff17049f2c4d.png	188890	image/png	http://localhost:9000/doscom-uploads/pengurus/20260608-cb913c09-eef7-4f19-a7fd-ff17049f2c4d.png	2026-06-08 12:37:44.471135	2026-06-08 12:37:44.471135
12	7	pengurus	Untitled design.png	pengurus/20260608-f5190948-04ef-49cb-8f1e-defd161bb098.png	325159	image/png	http://localhost:9000/doscom-uploads/pengurus/20260608-f5190948-04ef-49cb-8f1e-defd161bb098.png	2026-06-08 14:31:12.293886	2026-06-08 14:31:12.293886
13	7	pengurus	Untitled design.png	pengurus/20260608-5eabcb25-cbf7-4873-bdf1-327e9223e0ae.png	325159	image/png	http://localhost:9000/doscom-uploads/pengurus/20260608-5eabcb25-cbf7-4873-bdf1-327e9223e0ae.png	2026-06-08 14:51:05.842756	2026-06-08 14:51:05.842756
14	7	pengurus	Untitled design.png	pengurus/20260608-8f654d39-1145-42e1-bf0c-a8fa49c79f7e.png	325159	image/png	http://localhost:9000/doscom-uploads/pengurus/20260608-8f654d39-1145-42e1-bf0c-a8fa49c79f7e.png	2026-06-08 15:04:12.687595	2026-06-08 15:04:12.687595
15	7	pengurus	Untitled design.png	pengurus/20260608-9b4b33b9-7afa-4324-8f5b-d3fad314381a.png	325159	image/png	http://localhost:9000/doscom-uploads/pengurus/20260608-9b4b33b9-7afa-4324-8f5b-d3fad314381a.png	2026-06-08 15:08:15.161937	2026-06-08 15:08:15.161937
16	1	gallery	Classic Apple Pie - Eats Delightful.jpg	gallery/20260613-1afcfa77-9c57-4bbe-b7b0-5b7b3c96ee86.jpg	106625	image/jpeg	http://localhost:9000/doscom-uploads/gallery/20260613-1afcfa77-9c57-4bbe-b7b0-5b7b3c96ee86.jpg	2026-06-13 07:35:35.931001	2026-06-13 07:35:35.931001
17	28	work	Classic Apple Pie - Eats Delightful.jpg	work/20260614-31da7056-04c2-4a92-a3aa-8c8cf3079c57.jpg	106625	image/jpeg	http://localhost:9000/doscom-uploads/work/20260614-31da7056-04c2-4a92-a3aa-8c8cf3079c57.jpg	2026-06-14 04:19:54.34114	2026-06-14 04:19:54.34114
18	28	work	200.jpg	work/20260614-1c52adf6-cb08-4ce2-a970-34ae5e15cb72.jpg	20245	image/jpeg	http://localhost:9000/doscom-uploads/work/20260614-1c52adf6-cb08-4ce2-a970-34ae5e15cb72.jpg	2026-06-14 04:19:54.34114	2026-06-14 04:19:54.34114
19	28	work	Ice Bear Desktop Wallpaper.jpg	work/20260614-d79bb3f7-e34e-4d25-a2fa-40ca80aeea27.jpg	7673	image/jpeg	http://localhost:9000/doscom-uploads/work/20260614-d79bb3f7-e34e-4d25-a2fa-40ca80aeea27.jpg	2026-06-14 04:19:54.34114	2026-06-14 04:19:54.34114
20	28	work	Raining 🍃.jpg	work/20260614-9e1fe54e-4053-4c31-9926-f89ccd4c0f40.jpg	35144	image/jpeg	http://localhost:9000/doscom-uploads/work/20260614-9e1fe54e-4053-4c31-9926-f89ccd4c0f40.jpg	2026-06-14 04:19:54.34114	2026-06-14 04:19:54.34114
21	28	work	WhatsApp Image 2025-12-14 at 14.22.36.jpeg	work/20260614-fac6002b-94bb-407d-8fb2-9dc4cff1c6e6.jpeg	122236	image/jpeg	http://localhost:9000/doscom-uploads/work/20260614-fac6002b-94bb-407d-8fb2-9dc4cff1c6e6.jpeg	2026-06-14 04:19:54.34114	2026-06-14 04:19:54.34114
22	27	pengurus	Untitled design.png	pengurus/20260614-7ff05ff0-4a51-435a-b4c2-c9a5b5389033.png	325159	image/png	http://localhost:9000/doscom-uploads/pengurus/20260614-7ff05ff0-4a51-435a-b4c2-c9a5b5389033.png	2026-06-14 04:22:59.640297	2026-06-14 04:22:59.640297
38	27	work	Classic Apple Pie - Eats Delightful.jpg	work/20260614-46a69f73-15b9-4fe5-b43c-7dff90243d07.jpg	106625	image/jpeg	http://localhost:9000/doscom-uploads/work/20260614-46a69f73-15b9-4fe5-b43c-7dff90243d07.jpg	2026-06-14 08:26:28.432414	2026-06-14 08:26:28.432414
39	27	work	200.jpg	work/20260614-195f4ac6-f77a-4ace-b5db-936d13bc7682.jpg	20245	image/jpeg	http://localhost:9000/doscom-uploads/work/20260614-195f4ac6-f77a-4ace-b5db-936d13bc7682.jpg	2026-06-14 08:26:28.432414	2026-06-14 08:26:28.432414
40	27	work	Ice Bear Desktop Wallpaper.jpg	work/20260614-13372994-f7a6-4e1d-ac07-9421b05a34b6.jpg	7673	image/jpeg	http://localhost:9000/doscom-uploads/work/20260614-13372994-f7a6-4e1d-ac07-9421b05a34b6.jpg	2026-06-14 08:26:28.432414	2026-06-14 08:26:28.432414
41	27	work	Raining 🍃.jpg	work/20260614-9960a74d-e57a-46ef-8b6c-a310d35c4beb.jpg	35144	image/jpeg	http://localhost:9000/doscom-uploads/work/20260614-9960a74d-e57a-46ef-8b6c-a310d35c4beb.jpg	2026-06-14 08:26:28.432414	2026-06-14 08:26:28.432414
42	27	work	WhatsApp Image 2025-12-14 at 14.22.36.jpeg	work/20260614-49827106-b0d0-4abb-83f2-6874f1e34d04.jpeg	122236	image/jpeg	http://localhost:9000/doscom-uploads/work/20260614-49827106-b0d0-4abb-83f2-6874f1e34d04.jpeg	2026-06-14 08:26:28.432414	2026-06-14 08:26:28.432414
43	27	work	Classic Apple Pie - Eats Delightful.jpg	work/20260614-21b7f939-bf37-4202-98ab-33b9fca0e0c7.jpg	106625	image/jpeg	http://localhost:9000/doscom-uploads/work/20260614-21b7f939-bf37-4202-98ab-33b9fca0e0c7.jpg	2026-06-14 08:41:02.616857	2026-06-14 08:41:02.616857
44	27	work	200.jpg	work/20260614-1843b78d-b364-47ff-a7aa-61c44d614a55.jpg	20245	image/jpeg	http://localhost:9000/doscom-uploads/work/20260614-1843b78d-b364-47ff-a7aa-61c44d614a55.jpg	2026-06-14 08:41:02.616857	2026-06-14 08:41:02.616857
45	27	work	Ice Bear Desktop Wallpaper.jpg	work/20260614-f83da879-cac6-4132-b787-eae0b722387a.jpg	7673	image/jpeg	http://localhost:9000/doscom-uploads/work/20260614-f83da879-cac6-4132-b787-eae0b722387a.jpg	2026-06-14 08:41:02.616857	2026-06-14 08:41:02.616857
46	27	work	Raining 🍃.jpg	work/20260614-86610eef-c66a-49b7-a4ac-5127e748bfbd.jpg	35144	image/jpeg	http://localhost:9000/doscom-uploads/work/20260614-86610eef-c66a-49b7-a4ac-5127e748bfbd.jpg	2026-06-14 08:41:02.616857	2026-06-14 08:41:02.616857
47	1	work	Classic Apple Pie - Eats Delightful.jpg	work/20260614-929a5207-f8a7-4f4e-a977-7f1006eda381.jpg	106625	image/jpeg	http://localhost:9000/doscom-uploads/work/20260614-929a5207-f8a7-4f4e-a977-7f1006eda381.jpg	2026-06-14 09:21:50.080431	2026-06-14 09:21:50.080431
48	1	work	200.jpg	work/20260614-74459c1f-804a-4f41-b7ac-a4fa767ebf0a.jpg	20245	image/jpeg	http://localhost:9000/doscom-uploads/work/20260614-74459c1f-804a-4f41-b7ac-a4fa767ebf0a.jpg	2026-06-14 09:21:50.080431	2026-06-14 09:21:50.080431
49	1	work	Ice Bear Desktop Wallpaper.jpg	work/20260614-dd4d3c63-21c8-4bea-bfba-faf7d699a7ab.jpg	7673	image/jpeg	http://localhost:9000/doscom-uploads/work/20260614-dd4d3c63-21c8-4bea-bfba-faf7d699a7ab.jpg	2026-06-14 09:21:50.080431	2026-06-14 09:21:50.080431
50	1	work	Raining 🍃.jpg	work/20260614-98f24e7f-2934-4f95-bad1-27dffa3db2c4.jpg	35144	image/jpeg	http://localhost:9000/doscom-uploads/work/20260614-98f24e7f-2934-4f95-bad1-27dffa3db2c4.jpg	2026-06-14 09:21:50.080431	2026-06-14 09:21:50.080431
51	7	work	Classic Apple Pie - Eats Delightful.jpg	work/20260614-bf36757e-1b08-4a81-b3cd-4d4a5e0eac94.jpg	106625	image/jpeg	http://localhost:9000/doscom-uploads/work/20260614-bf36757e-1b08-4a81-b3cd-4d4a5e0eac94.jpg	2026-06-14 12:01:27.861149	2026-06-14 12:01:27.861149
52	7	work	200.jpg	work/20260614-080173f6-205c-4918-b5a3-2ad77f96701c.jpg	20245	image/jpeg	http://localhost:9000/doscom-uploads/work/20260614-080173f6-205c-4918-b5a3-2ad77f96701c.jpg	2026-06-14 12:01:27.861149	2026-06-14 12:01:27.861149
53	7	work	Ice Bear Desktop Wallpaper.jpg	work/20260614-269b7d38-c945-4201-8496-ff861955844a.jpg	7673	image/jpeg	http://localhost:9000/doscom-uploads/work/20260614-269b7d38-c945-4201-8496-ff861955844a.jpg	2026-06-14 12:01:27.861149	2026-06-14 12:01:27.861149
54	7	work	Raining 🍃.jpg	work/20260614-0521164d-02be-4915-819b-ab542a26c0e3.jpg	35144	image/jpeg	http://localhost:9000/doscom-uploads/work/20260614-0521164d-02be-4915-819b-ab542a26c0e3.jpg	2026-06-14 12:01:27.861149	2026-06-14 12:01:27.861149
55	7	work	Classic Apple Pie - Eats Delightful.jpg	work/20260615-8ae17b7d-5da0-4554-91de-6d3b2251b051.jpg	106625	image/jpeg	http://localhost:9000/doscom-uploads/work/20260615-8ae17b7d-5da0-4554-91de-6d3b2251b051.jpg	2026-06-15 10:16:06.272059	2026-06-15 10:16:06.272059
56	7	work	200.jpg	work/20260615-89ff3d49-e575-4e18-b778-6928737a69e9.jpg	20245	image/jpeg	http://localhost:9000/doscom-uploads/work/20260615-89ff3d49-e575-4e18-b778-6928737a69e9.jpg	2026-06-15 10:16:06.272059	2026-06-15 10:16:06.272059
57	7	work	Ice Bear Desktop Wallpaper.jpg	work/20260615-5bc6b14f-0a79-40b6-b967-795e9fa27626.jpg	7673	image/jpeg	http://localhost:9000/doscom-uploads/work/20260615-5bc6b14f-0a79-40b6-b967-795e9fa27626.jpg	2026-06-15 10:16:06.272059	2026-06-15 10:16:06.272059
58	7	work	Raining 🍃.jpg	work/20260615-a60fe867-2e9b-459c-b953-8434fc72d93e.jpg	35144	image/jpeg	http://localhost:9000/doscom-uploads/work/20260615-a60fe867-2e9b-459c-b953-8434fc72d93e.jpg	2026-06-15 10:16:06.272059	2026-06-15 10:16:06.272059
59	7	work	Classic Apple Pie - Eats Delightful.jpg	work/20260615-06507c1b-a6b8-4e85-93d1-d42c60a43055.jpg	106625	image/jpeg	http://localhost:9000/doscom-uploads/work/20260615-06507c1b-a6b8-4e85-93d1-d42c60a43055.jpg	2026-06-15 10:17:21.588113	2026-06-15 10:17:21.588113
60	7	work	200.jpg	work/20260615-ec5eebf8-00b5-4f0e-b7af-bab3d7f5e701.jpg	20245	image/jpeg	http://localhost:9000/doscom-uploads/work/20260615-ec5eebf8-00b5-4f0e-b7af-bab3d7f5e701.jpg	2026-06-15 10:17:21.588113	2026-06-15 10:17:21.588113
61	7	work	Ice Bear Desktop Wallpaper.jpg	work/20260615-2442358b-c133-4e8f-acf7-80e3d0039704.jpg	7673	image/jpeg	http://localhost:9000/doscom-uploads/work/20260615-2442358b-c133-4e8f-acf7-80e3d0039704.jpg	2026-06-15 10:17:21.588113	2026-06-15 10:17:21.588113
62	7	work	Raining 🍃.jpg	work/20260615-e2307561-cf9a-4ba5-91d5-a5bdc17924b6.jpg	35144	image/jpeg	http://localhost:9000/doscom-uploads/work/20260615-e2307561-cf9a-4ba5-91d5-a5bdc17924b6.jpg	2026-06-15 10:17:21.588113	2026-06-15 10:17:21.588113
63	7	work	Classic Apple Pie - Eats Delightful.jpg	work/20260615-8c5d553d-198b-4cbb-81fa-276daa8ade44.jpg	106625	image/jpeg	http://localhost:9000/doscom-uploads/work/20260615-8c5d553d-198b-4cbb-81fa-276daa8ade44.jpg	2026-06-15 14:56:12.821329	2026-06-15 14:56:12.821329
64	7	work	200.jpg	work/20260615-3a11f0df-de5b-496a-9600-cadbe555bb6b.jpg	20245	image/jpeg	http://localhost:9000/doscom-uploads/work/20260615-3a11f0df-de5b-496a-9600-cadbe555bb6b.jpg	2026-06-15 14:56:12.821329	2026-06-15 14:56:12.821329
65	7	work	Classic Apple Pie - Eats Delightful.jpg	work/20260615-ba481294-c570-440c-ab75-eae5b8c0b142.jpg	106625	image/jpeg	http://localhost:9000/doscom-uploads/work/20260615-ba481294-c570-440c-ab75-eae5b8c0b142.jpg	2026-06-15 15:05:10.731194	2026-06-15 15:05:10.731194
66	7	work	200.jpg	work/20260615-91d596f3-bec7-4210-9f1f-27004565bd5a.jpg	20245	image/jpeg	http://localhost:9000/doscom-uploads/work/20260615-91d596f3-bec7-4210-9f1f-27004565bd5a.jpg	2026-06-15 15:05:10.731194	2026-06-15 15:05:10.731194
67	7	work	Classic Apple Pie - Eats Delightful.jpg	work/20260615-0fa4e2fa-18e1-4951-bbf8-2ef7b5005066.jpg	106625	image/jpeg	http://localhost:9000/doscom-uploads/work/20260615-0fa4e2fa-18e1-4951-bbf8-2ef7b5005066.jpg	2026-06-15 16:39:45.700175	2026-06-15 16:39:45.700175
68	7	work	200.jpg	work/20260615-2cd99c98-4396-43f0-9c94-2a21d5704dec.jpg	20245	image/jpeg	http://localhost:9000/doscom-uploads/work/20260615-2cd99c98-4396-43f0-9c94-2a21d5704dec.jpg	2026-06-15 16:39:45.700175	2026-06-15 16:39:45.700175
69	7	work	Ice Bear Desktop Wallpaper.jpg	work/20260615-fd91b61c-8cde-4bee-95fb-ba388cc58cda.jpg	7673	image/jpeg	http://localhost:9000/doscom-uploads/work/20260615-fd91b61c-8cde-4bee-95fb-ba388cc58cda.jpg	2026-06-15 16:39:45.700175	2026-06-15 16:39:45.700175
70	7	work	Raining 🍃.jpg	work/20260615-582656db-51bb-49dd-9cbd-ec72031e38ba.jpg	35144	image/jpeg	http://localhost:9000/doscom-uploads/work/20260615-582656db-51bb-49dd-9cbd-ec72031e38ba.jpg	2026-06-15 16:39:45.700175	2026-06-15 16:39:45.700175
71	29	pengurus	26-Photoroom.png	pengurus/20260615-5cf14ad8-21c1-4ed5-8519-91a8dfd0b08d.png	188890	image/png	http://localhost:9000/doscom-uploads/pengurus/20260615-5cf14ad8-21c1-4ed5-8519-91a8dfd0b08d.png	2026-06-15 17:15:23.342069	2026-06-15 17:15:23.342069
72	7	pengurus	Untitled design.png	pengurus/20260616-d40a9528-c108-4b3b-b980-cdebf7aea8a9.png	325159	image/png	http://localhost:9000/doscom-uploads/pengurus/20260616-d40a9528-c108-4b3b-b980-cdebf7aea8a9.png	2026-06-16 06:02:01.504647	2026-06-16 06:02:01.504647
73	7	pengurus	Untitled design.png	pengurus/20260616-b8ac736c-d26a-4804-8b56-8d760ea27f08.png	325159	image/png	http://localhost:9000/doscom-uploads/pengurus/20260616-b8ac736c-d26a-4804-8b56-8d760ea27f08.png	2026-06-16 06:02:12.954659	2026-06-16 06:02:12.954659
\.


--
-- Data for Name: gallery; Type: TABLE DATA; Schema: public; Owner: iko
--

COPY public.gallery (id, id_users, file_upload_id, gallery_name, gallery_type, description, event_date, file_url, created_at, updated_at) FROM stdin;
5	7	5	foto profil pengurus	pengurus	foto identitas diri yang mewakili pengurus doscom	2026-05-29 00:00:00		2026-05-29 11:46:52.580195	2026-05-29 11:46:52.580195
4	\N	4	foto profil pengurus	pengurus	foto identitas diri yang mewakili pengurus doscom	2026-05-28 00:00:00		2026-05-28 19:32:44.634636	2026-05-28 19:32:44.634636
6	21	6	foto profil pengurus	pengurus	foto identitas diri yang mewakili pengurus doscom	2026-06-08 00:00:00		2026-06-08 11:22:09.358392	2026-06-08 11:22:09.358392
7	21	7	foto profil pengurus	pengurus	foto identitas diri yang mewakili pengurus doscom	2026-06-08 00:00:00		2026-06-08 11:23:05.264516	2026-06-08 11:23:05.264516
8	21	8	foto profil pengurus	pengurus	foto identitas diri yang mewakili pengurus doscom	2026-06-08 00:00:00		2026-06-08 11:31:09.9314	2026-06-08 11:31:09.9314
9	21	9	foto profil pengurus	pengurus	foto identitas diri yang mewakili pengurus doscom	2026-06-08 00:00:00		2026-06-08 12:23:11.036266	2026-06-08 12:23:11.036266
10	21	10	foto profil pengurus	pengurus	foto identitas diri yang mewakili pengurus doscom	2026-06-08 00:00:00		2026-06-08 12:24:31.899466	2026-06-08 12:24:31.899466
11	21	11	foto profil pengurus	pengurus	foto identitas diri yang mewakili pengurus doscom	2026-06-08 00:00:00		2026-06-08 12:37:44.474586	2026-06-08 12:37:44.474586
14	26	14	foto profil pengurus	pengurus	foto identitas diri yang mewakili pengurus doscom	2026-06-08 00:00:00		2026-06-08 15:04:12.691577	2026-06-08 15:04:12.691578
15	3	15	foto profil pengurus	pengurus	foto identitas diri yang mewakili pengurus doscom	2026-06-08 00:00:00		2026-06-08 15:08:15.16322	2026-06-08 15:08:15.16322
16	1	16	foto project pie numpy	proker	ini adalah foto sebuah makanan yaitu kue pie rasanya yag biasa saja	2025-12-12 00:00:00	http://localhost:9000/doscom-uploads/gallery/20260613-1afcfa77-9c57-4bbe-b7b0-5b7b3c96ee86.jpg	2026-06-13 07:35:35.935976	2026-06-13 07:35:35.935976
17	28	17	foto untuk work dengan judulcolor pallete retro	work	foto untuk kepentingan work	2026-06-14 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-31da7056-04c2-4a92-a3aa-8c8cf3079c57.jpg	2026-06-14 04:19:54.352791	2026-06-14 04:19:54.352791
18	28	18	foto untuk work dengan judulcolor pallete retro	work	foto untuk kepentingan work	2026-06-14 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-1c52adf6-cb08-4ce2-a970-34ae5e15cb72.jpg	2026-06-14 04:19:54.352793	2026-06-14 04:19:54.352793
19	28	19	foto untuk work dengan judulcolor pallete retro	work	foto untuk kepentingan work	2026-06-14 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-d79bb3f7-e34e-4d25-a2fa-40ca80aeea27.jpg	2026-06-14 04:19:54.352794	2026-06-14 04:19:54.352795
20	28	20	foto untuk work dengan judulcolor pallete retro	work	foto untuk kepentingan work	2026-06-14 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-9e1fe54e-4053-4c31-9926-f89ccd4c0f40.jpg	2026-06-14 04:19:54.352795	2026-06-14 04:19:54.352795
21	28	21	foto untuk work dengan judulcolor pallete retro	work	foto untuk kepentingan work	2026-06-14 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-fac6002b-94bb-407d-8fb2-9dc4cff1c6e6.jpeg	2026-06-14 04:19:54.352796	2026-06-14 04:19:54.352796
22	28	22	foto profil pengurus	pengurus	foto identitas diri yang mewakili pengurus doscom	2026-06-14 00:00:00		2026-06-14 04:22:59.643266	2026-06-14 04:22:59.643266
23	27	38	foto untuk work dengan judulcolor pallete retro	work	foto untuk kepentingan work	2026-06-14 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-46a69f73-15b9-4fe5-b43c-7dff90243d07.jpg	2026-06-14 08:26:28.435494	2026-06-14 08:26:28.435494
24	27	39	foto untuk work dengan judulcolor pallete retro	work	foto untuk kepentingan work	2026-06-14 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-195f4ac6-f77a-4ace-b5db-936d13bc7682.jpg	2026-06-14 08:26:28.435494	2026-06-14 08:26:28.435494
25	27	40	foto untuk work dengan judulcolor pallete retro	work	foto untuk kepentingan work	2026-06-14 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-13372994-f7a6-4e1d-ac07-9421b05a34b6.jpg	2026-06-14 08:26:28.435494	2026-06-14 08:26:28.435495
26	27	41	foto untuk work dengan judulcolor pallete retro	work	foto untuk kepentingan work	2026-06-14 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-9960a74d-e57a-46ef-8b6c-a310d35c4beb.jpg	2026-06-14 08:26:28.435495	2026-06-14 08:26:28.435495
27	27	42	foto untuk work dengan judulcolor pallete retro	work	foto untuk kepentingan work	2026-06-14 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-49827106-b0d0-4abb-83f2-6874f1e34d04.jpeg	2026-06-14 08:26:28.435495	2026-06-14 08:26:28.435495
28	27	43	foto untuk work dengan judulweb pentester	work	foto untuk kepentingan work	2026-06-14 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-21b7f939-bf37-4202-98ab-33b9fca0e0c7.jpg	2026-06-14 08:41:02.618507	2026-06-14 08:41:02.618507
29	27	44	foto untuk work dengan judulweb pentester	work	foto untuk kepentingan work	2026-06-14 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-1843b78d-b364-47ff-a7aa-61c44d614a55.jpg	2026-06-14 08:41:02.618507	2026-06-14 08:41:02.618507
30	27	45	foto untuk work dengan judulweb pentester	work	foto untuk kepentingan work	2026-06-14 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-f83da879-cac6-4132-b787-eae0b722387a.jpg	2026-06-14 08:41:02.618507	2026-06-14 08:41:02.618507
31	27	46	foto untuk work dengan judulweb pentester	work	foto untuk kepentingan work	2026-06-14 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-86610eef-c66a-49b7-a4ac-5127e748bfbd.jpg	2026-06-14 08:41:02.618507	2026-06-14 08:41:02.618507
32	1	47	foto untuk work dengan judulsicon tools	work	foto untuk kepentingan work	2026-06-14 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-929a5207-f8a7-4f4e-a977-7f1006eda381.jpg	2026-06-14 09:21:50.082946	2026-06-14 09:21:50.082946
33	1	48	foto untuk work dengan judulsicon tools	work	foto untuk kepentingan work	2026-06-14 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-74459c1f-804a-4f41-b7ac-a4fa767ebf0a.jpg	2026-06-14 09:21:50.082947	2026-06-14 09:21:50.082947
34	1	49	foto untuk work dengan judulsicon tools	work	foto untuk kepentingan work	2026-06-14 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-dd4d3c63-21c8-4bea-bfba-faf7d699a7ab.jpg	2026-06-14 09:21:50.082947	2026-06-14 09:21:50.082947
35	1	50	foto untuk work dengan judulsicon tools	work	foto untuk kepentingan work	2026-06-14 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-98f24e7f-2934-4f95-bad1-27dffa3db2c4.jpg	2026-06-14 09:21:50.082947	2026-06-14 09:21:50.082947
36	7	51	foto untuk work dengan juduld-form web	work	foto untuk kepentingan work	2026-06-14 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-bf36757e-1b08-4a81-b3cd-4d4a5e0eac94.jpg	2026-06-14 12:01:27.87118	2026-06-14 12:01:27.871181
37	7	52	foto untuk work dengan juduld-form web	work	foto untuk kepentingan work	2026-06-14 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-080173f6-205c-4918-b5a3-2ad77f96701c.jpg	2026-06-14 12:01:27.871182	2026-06-14 12:01:27.871182
38	7	53	foto untuk work dengan juduld-form web	work	foto untuk kepentingan work	2026-06-14 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-269b7d38-c945-4201-8496-ff861955844a.jpg	2026-06-14 12:01:27.871182	2026-06-14 12:01:27.871182
39	7	54	foto untuk work dengan juduld-form web	work	foto untuk kepentingan work	2026-06-14 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-0521164d-02be-4915-819b-ab542a26c0e3.jpg	2026-06-14 12:01:27.871183	2026-06-14 12:01:27.871183
40	7	55	foto untuk work dengan judulweb doscom university	work	foto untuk kepentingan work	2026-06-15 00:00:00	http://localhost:9000/doscom-uploads/work/20260615-8ae17b7d-5da0-4554-91de-6d3b2251b051.jpg	2026-06-15 10:16:06.277041	2026-06-15 10:16:06.277041
41	7	56	foto untuk work dengan judulweb doscom university	work	foto untuk kepentingan work	2026-06-15 00:00:00	http://localhost:9000/doscom-uploads/work/20260615-89ff3d49-e575-4e18-b778-6928737a69e9.jpg	2026-06-15 10:16:06.277041	2026-06-15 10:16:06.277041
42	7	57	foto untuk work dengan judulweb doscom university	work	foto untuk kepentingan work	2026-06-15 00:00:00	http://localhost:9000/doscom-uploads/work/20260615-5bc6b14f-0a79-40b6-b967-795e9fa27626.jpg	2026-06-15 10:16:06.277041	2026-06-15 10:16:06.277042
43	7	58	foto untuk work dengan judulweb doscom university	work	foto untuk kepentingan work	2026-06-15 00:00:00	http://localhost:9000/doscom-uploads/work/20260615-a60fe867-2e9b-459c-b953-8434fc72d93e.jpg	2026-06-15 10:16:06.277042	2026-06-15 10:16:06.277042
44	7	59	foto untuk work dengan judulweb doscom university	work	foto untuk kepentingan work	2026-06-15 00:00:00	http://localhost:9000/doscom-uploads/work/20260615-06507c1b-a6b8-4e85-93d1-d42c60a43055.jpg	2026-06-15 10:17:21.590172	2026-06-15 10:17:21.590172
45	7	60	foto untuk work dengan judulweb doscom university	work	foto untuk kepentingan work	2026-06-15 00:00:00	http://localhost:9000/doscom-uploads/work/20260615-ec5eebf8-00b5-4f0e-b7af-bab3d7f5e701.jpg	2026-06-15 10:17:21.590173	2026-06-15 10:17:21.590173
46	7	61	foto untuk work dengan judulweb doscom university	work	foto untuk kepentingan work	2026-06-15 00:00:00	http://localhost:9000/doscom-uploads/work/20260615-2442358b-c133-4e8f-acf7-80e3d0039704.jpg	2026-06-15 10:17:21.590173	2026-06-15 10:17:21.590173
47	7	62	foto untuk work dengan judulweb doscom university	work	foto untuk kepentingan work	2026-06-15 00:00:00	http://localhost:9000/doscom-uploads/work/20260615-e2307561-cf9a-4ba5-91d5-a5bdc17924b6.jpg	2026-06-15 10:17:21.590173	2026-06-15 10:17:21.590173
48	7	63	foto untuk work dengan judulricesource	work	foto untuk kepentingan work	2026-06-15 00:00:00	http://localhost:9000/doscom-uploads/work/20260615-8c5d553d-198b-4cbb-81fa-276daa8ade44.jpg	2026-06-15 14:56:12.826319	2026-06-15 14:56:12.826319
49	7	64	foto untuk work dengan judulricesource	work	foto untuk kepentingan work	2026-06-15 00:00:00	http://localhost:9000/doscom-uploads/work/20260615-3a11f0df-de5b-496a-9600-cadbe555bb6b.jpg	2026-06-15 14:56:12.826319	2026-06-15 14:56:12.826319
50	7	65	foto untuk work dengan judulricesource	work	foto untuk kepentingan work	2026-06-15 00:00:00	http://localhost:9000/doscom-uploads/work/20260615-ba481294-c570-440c-ab75-eae5b8c0b142.jpg	2026-06-15 15:05:10.733548	2026-06-15 15:05:10.733548
51	7	66	foto untuk work dengan judulricesource	work	foto untuk kepentingan work	2026-06-15 00:00:00	http://localhost:9000/doscom-uploads/work/20260615-91d596f3-bec7-4210-9f1f-27004565bd5a.jpg	2026-06-15 15:05:10.733548	2026-06-15 15:05:10.733548
52	7	67	foto untuk work dengan judulricesource	work	foto untuk kepentingan work	2026-06-15 00:00:00	http://localhost:9000/doscom-uploads/work/20260615-0fa4e2fa-18e1-4951-bbf8-2ef7b5005066.jpg	2026-06-15 16:39:45.703287	2026-06-15 16:39:45.703287
53	7	68	foto untuk work dengan judulricesource	work	foto untuk kepentingan work	2026-06-15 00:00:00	http://localhost:9000/doscom-uploads/work/20260615-2cd99c98-4396-43f0-9c94-2a21d5704dec.jpg	2026-06-15 16:39:45.703289	2026-06-15 16:39:45.703289
54	7	69	foto untuk work dengan judulricesource	work	foto untuk kepentingan work	2026-06-15 00:00:00	http://localhost:9000/doscom-uploads/work/20260615-fd91b61c-8cde-4bee-95fb-ba388cc58cda.jpg	2026-06-15 16:39:45.70329	2026-06-15 16:39:45.70329
55	7	70	foto untuk work dengan judulricesource	work	foto untuk kepentingan work	2026-06-15 00:00:00	http://localhost:9000/doscom-uploads/work/20260615-582656db-51bb-49dd-9cbd-ec72031e38ba.jpg	2026-06-15 16:39:45.703291	2026-06-15 16:39:45.703291
56	29	71	foto profil pengurus	pengurus	foto identitas diri yang mewakili pengurus doscom	2026-06-15 00:00:00		2026-06-15 17:15:23.344549	2026-06-15 17:15:23.344549
57	25	72	foto profil pengurus	pengurus	foto identitas diri yang mewakili pengurus doscom	2026-06-16 00:00:00		2026-06-16 06:02:01.509691	2026-06-16 06:02:01.509691
58	25	73	foto profil pengurus	pengurus	foto identitas diri yang mewakili pengurus doscom	2026-06-16 00:00:00		2026-06-16 06:02:12.956163	2026-06-16 06:02:12.956163
\.


--
-- Data for Name: history_photos; Type: TABLE DATA; Schema: public; Owner: iko
--

COPY public.history_photos (id, id_history, imager_url, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: history_timeline; Type: TABLE DATA; Schema: public; Owner: iko
--

COPY public.history_timeline (id, id_author, title, year, description, display_order, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: pengurus; Type: TABLE DATA; Schema: public; Owner: iko
--

COPY public.pengurus (id, id_user, photo_url, email, divisi, name, "position", start_periode_year, end_periode_year, created_at, updated_at) FROM stdin;
1	2		jaringan@doscom.org	jaringan	muhammad daniyal haq	KoorJaringan	2023	2024	2026-05-27 06:26:32.822029	2026-05-27 06:26:32.816034
2	3		anggotajaringan@doscom.org	jaringan	fujikawa shinici	JaringanAng	2024	2027	2026-05-27 06:26:32.827411	2026-05-27 06:26:32.816034
10	7	http://localhost:9000/doscom-uploads/pengurus/20260529-5506a862-30cd-4749-8d82-13fc18e48bab.png	anakbaikkok@gmail.com	pemro	rico andre pratama	PemroAng	2024	2026	2026-05-29 11:46:52.593418	2026-05-29 11:46:52.593646
16	26	http://localhost:9000/doscom-uploads/pengurus/20260608-8f654d39-1145-42e1-bf0c-a8fa49c79f7e.png	arrelsukafurry@gmail.com	pemro	danendra farrel	PemroAng	2025	2027	2026-06-08 15:04:12.696589	2026-06-08 15:04:12.696649
17	3	http://localhost:9000/doscom-uploads/pengurus/20260608-9b4b33b9-7afa-4324-8f5b-d3fad314381a.png	furrysukabre@gmail.com	pemro	brenendra femboy	PemroAng	2025	2027	2026-06-08 15:08:15.164866	2026-06-08 15:08:15.164896
3	4		brewongsangarwak@gmail.com	medcrev	dhion oppa	KoorMedcrev	2023	2024	2026-05-27 06:26:32.831245	2026-06-08 15:12:52.265651
18	28	http://localhost:9000/doscom-uploads/pengurus/20260614-7ff05ff0-4a51-435a-b4c2-c9a5b5389033.png	akukacamata@gmail.com	medcrev	zovan yang z	MedcrevAng	2025	2027	2026-06-14 04:22:59.648499	2026-06-14 04:22:59.64861
20	25	http://localhost:9000/doscom-uploads/pengurus/20260616-b8ac736c-d26a-4804-8b56-8d760ea27f08.png	wongirengkeren@gmail.com	pemro	zovan yang z	PemroAng	2025	2027	2026-06-16 06:02:12.958547	2026-06-16 06:02:12.958606
\.


--
-- Data for Name: pengurus_sosmed; Type: TABLE DATA; Schema: public; Owner: iko
--

COPY public.pengurus_sosmed (id, pengurus_id, platform, username, url, is_primary, created_at, updated_at) FROM stdin;
1	1	instagram	dhioonnn	https://www.instagram.com/dhioonnn/	t	2026-05-27 06:26:32.837506	2026-05-27 06:26:32.837506
2	1	linkedin	dhion-nur-damanhuri-2bb863275	https://www.linkedin.com/in/dhion-nur-damanhuri-2bb863275/	f	2026-05-27 06:26:32.837506	2026-05-27 06:26:32.837506
3	1	github	IKOPOO	https://github.com/IKOPOO	f	2026-05-27 06:26:32.837506	2026-05-27 06:26:32.837506
4	2	instagram	dhioonnn	https://www.instagram.com/dhioonnn/	t	2026-05-27 06:26:32.841828	2026-05-27 06:26:32.841828
5	2	linkedin	dhion-nur-damanhuri-2bb863275	https://www.linkedin.com/in/dhion-nur-damanhuri-2bb863275/	f	2026-05-27 06:26:32.841828	2026-05-27 06:26:32.841828
6	2	github	IKOPOO	https://github.com/IKOPOO	f	2026-05-27 06:26:32.841828	2026-05-27 06:26:32.841828
28	10	instagram	ricoandreprtm	https://www.instagram.com/ricoandreprtm/	t	2026-05-29 11:46:52.599302	2026-05-29 11:46:52.599302
29	10	linkedin	ricoandrepratama	https://www.linkedin.com/in/ricoandrepratama/	f	2026-05-29 11:46:52.599302	2026-05-29 11:46:52.599302
30	10	github	IKOPOO	https://github.com/IKOPOO	f	2026-05-29 11:46:52.599302	2026-05-29 11:46:52.599302
38	16	instagram	ricoandreprtm	https://www.instagram.com/ricoandreprtm/	t	2026-06-08 15:04:12.699929	2026-06-08 15:04:12.699929
39	16	linkedin	ricoandrepratama	https://www.linkedin.com/in/ricoandrepratama/	f	2026-06-08 15:04:12.699929	2026-06-08 15:04:12.699929
40	16	github	IKOPOO	https://github.com/IKOPOO	f	2026-06-08 15:04:12.699929	2026-06-08 15:04:12.699929
41	17	instagram	ricoandreprtm	https://www.instagram.com/ricoandreprtm/	t	2026-06-08 15:08:15.16601	2026-06-08 15:08:15.16601
42	17	linkedin	ricoandrepratama	https://www.linkedin.com/in/ricoandrepratama/	f	2026-06-08 15:08:15.16601	2026-06-08 15:08:15.16601
43	17	github	IKOPOO	https://github.com/IKOPOO	f	2026-06-08 15:08:15.16601	2026-06-08 15:08:15.16601
44	3	linkedin	ricoandrepratama	https://www.linkedin.com/in/ricoandrepratama/	t	2026-06-08 15:12:52.268492	2026-06-08 15:12:52.268492
45	3	github	IKOPOO	https://github.com/IKOPOO	f	2026-06-08 15:12:52.268492	2026-06-08 15:12:52.268492
46	18	instagram	ricoandreprtm	https://www.instagram.com/ricoandreprtm/	t	2026-06-14 04:22:59.650988	2026-06-14 04:22:59.650988
47	18	linkedin	ricoandrepratama	https://www.linkedin.com/in/ricoandrepratama/	f	2026-06-14 04:22:59.650988	2026-06-14 04:22:59.650988
48	18	github	IKOPOO	https://github.com/IKOPOO	f	2026-06-14 04:22:59.650988	2026-06-14 04:22:59.650988
52	20	instagram	ricoandreprtm	https://www.instagram.com/ricoandreprtm/	t	2026-06-16 06:02:12.960374	2026-06-16 06:02:12.960374
53	20	linkedin	ricoandrepratama	https://www.linkedin.com/in/ricoandrepratama/	f	2026-06-16 06:02:12.960374	2026-06-16 06:02:12.960374
54	20	github	IKOPOO	https://github.com/IKOPOO	f	2026-06-16 06:02:12.960374	2026-06-16 06:02:12.960374
\.


--
-- Data for Name: refresh_token; Type: TABLE DATA; Schema: public; Owner: iko
--

COPY public.refresh_token (id, user_id, token, expires, created_at, updated_at) FROM stdin;
89	7	548a3449c436e98288c2e7021494792720b13340ad56666237ef9790c78d86e0	2026-06-21 05:52:29.852542	2026-06-16 05:52:29.852543	2026-06-16 05:52:29.852543
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: iko
--

COPY public.schema_migrations (version, dirty) FROM stdin;
12	f
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: iko
--

COPY public.users (id, email, username, role, password, full_name, created_at, updated_at) FROM stdin;
1	superadmin@doscom.org	superadmin	SuperAdmin	$argon2id$v=19$m=65536,t=1,p=16$mzYD5xjNopZcQEE2nQndtQ$XB/0oapqhWjSrKIfNBHPsccvb4TtZHCxLYPToh8+ths	Super Admin Doscom	2026-05-27 06:26:32.764833	2026-05-27 06:26:32.764833
2	jaringan@doscom.org	koor jaringan	KoorJaringan	$argon2id$v=19$m=65536,t=1,p=16$mzYD5xjNopZcQEE2nQndtQ$XB/0oapqhWjSrKIfNBHPsccvb4TtZHCxLYPToh8+ths	muhammad daniyal haq	2026-05-27 06:26:32.764833	2026-05-27 06:26:32.764833
3	anggotajaringan@doscom.org	anggota jaringan	jaringanAnggota	$argon2id$v=19$m=65536,t=1,p=16$mzYD5xjNopZcQEE2nQndtQ$XB/0oapqhWjSrKIfNBHPsccvb4TtZHCxLYPToh8+ths	fujikawa nippon	2026-05-27 06:26:32.764833	2026-05-27 06:26:32.764833
4	medcrev@doscom.org	dhion sangar boss	KoorMedcrev	$argon2id$v=19$m=65536,t=1,p=16$mzYD5xjNopZcQEE2nQndtQ$XB/0oapqhWjSrKIfNBHPsccvb4TtZHCxLYPToh8+ths	dhion oppa	2026-05-27 06:26:32.764833	2026-05-27 06:26:32.764833
5	anggotamedcrev@doscom.org	dhion tapi anggota	medcrevAnggota	$argon2id$v=19$m=65536,t=1,p=16$mzYD5xjNopZcQEE2nQndtQ$XB/0oapqhWjSrKIfNBHPsccvb4TtZHCxLYPToh8+ths	dhion keren	2026-05-27 06:26:32.764833	2026-05-27 06:26:32.764833
6	anotherSuperadmin@doscom.org	superadmin_anotherSuperadmin	SuperAdmin	$argon2id$v=19$m=65536,t=1,p=16$3tut/fg1ZG48VuSCY05PzA$6BlxuKzRw2UXawZCSMXzdpmOkCO/h3pBzdSKiC+5qa8	superadmin	2026-05-27 06:28:31.740007	2026-05-27 06:28:31.740007
25	jomokngawi@gmail.com	alip_jomokngawi	pemroAnggota	$argon2id$v=19$m=65536,t=1,p=16$rcr75xwK9ccUBgR1gYWOrg$f5ttSuveWG5+Y8Xkg5EFAVEPtlWLAlI9qybrvJRSeOE	alip padang	2026-06-04 16:19:33.455116	2026-06-04 16:19:33.455116
9	alipketum@gmail.com	husnul_alipketum	pemroAnggota	$argon2id$v=19$m=65536,t=1,p=16$l+1ZjT4FwGAgyNmG0GkcVg$hrBHIUp9XeHyVmEYodBiZZ2I3p1273buXjIhh6VuL8c	husnul fikri averus	2026-05-27 06:48:35.811607	2026-05-27 06:48:35.811607
10	zapto@gmail.com	sapto_zapto	pemroAnggota	$argon2id$v=19$m=65536,t=1,p=16$fcxBd0t1AwrLkCxIrxjDxQ$J6fXzB3nv0c+tTLcUFmWJI0331K1ccjmEWRbkxT5xTo	sapto gusti agung	2026-05-27 12:42:00.155684	2026-05-27 12:42:00.155684
17	ambatukamrusdi@gmail.com	nippon_ambatukamrusdi	pemroAnggota	$argon2id$v=19$m=65536,t=1,p=16$WWTflokZA8grGiScKh1sFg$rxuKf55w10S6wPOORhBjeITToG184v69EPvaKVY/KDw	nippon ngawi	2026-05-29 16:11:14.754827	2026-05-29 16:11:14.754827
20	niggersuperadmin@doscom.org	nigga_niggersuperadmin	SuperAdmin	$argon2id$v=19$m=65536,t=1,p=16$8QforhVn2AgxLMbI+WHhxg$cE0PkQ+3h+0Wp1qkbGJnoZAuK9n+eev/7mG8FoU0OK4	nigga	2026-05-31 15:56:08.892016	2026-05-31 15:56:08.892016
11	wongsangar@gmail.com	fajar ganteng cihuyy	pemroAnggota	$argon2id$v=19$m=65536,t=1,p=16$crRZdOChRJM2ZUFc9Qn0jw$kQmiYrqkSEXf1O5xDHBT8Vk85dNVxYuYDCLSqmwDUsI	fajar aziz zisss	2026-05-27 15:06:09.677945	2026-05-31 20:24:12.833505
7	owentheowl@gmail.com	rico_ikotheturtle	KoorPemro	$argon2id$v=19$m=65536,t=1,p=16$rVG+AaDbS3STrGlmQRcYrQ$JiGBq06DNzyPUIEk+8ucvyNe0Re/IOIOwg4cIPR3gBc	rico andre pratama	2026-05-27 06:29:17.666047	2026-05-31 20:31:46.847772
22	superadminketiga@doscom.org	superadmin_superadminketiga	SuperAdmin	$argon2id$v=19$m=65536,t=1,p=16$pikKmGOe0gYmt6Ny+6SH0g$wg4QA8CZrcudIALRRiBWlblZWH7PA4smnNdM1l+/S3E	superadmin	2026-06-04 15:10:23.601298	2026-06-04 15:10:23.601298
19	ngawiambatukam@gmail.com	kevin_ngawiambatukam	pemroAnggota	$argon2id$v=19$m=65536,t=1,p=16$ZyBqQYSwWfdb3qF7H39OOQ$qbpBeQUqu8H8iRzYcX3Le89BTf2C/O0YtQTGjn0mCXY	kevin stuart	2026-05-30 06:02:36.563348	2026-06-04 15:20:00.656352
21	hitlerwongnganjuk@gmail.com	erika	pemroAnggota	$argon2id$v=19$m=65536,t=1,p=16$gVSCroCLAZ1tGjsedBfyMg$5j2NpWYu10awSputOUvpd3oeLPUvmM73XkFKTby3pwY	hail hitler nganjuk	2026-05-31 20:33:52.425027	2026-06-04 16:34:53.180185
26	rusdingawi@gmail.com	danendra_rusdingawi	pemroAnggota	$argon2id$v=19$m=65536,t=1,p=16$wyV+DPBH/viLQ4DDt6TfXg$hNkhVNAV8AvJkEy85wfTTMSn07fgunFiOei7/BaEs2E	danendra	2026-06-08 14:27:34.971347	2026-06-08 14:27:34.971347
27	dhionkoor@gmail.com	dhion_dhionkoor	KoorMedcrev	$argon2id$v=19$m=65536,t=1,p=16$1z0RTXcw1puVPABeZpyTlg$CwxW3oQrUaJpzvvF2QQykyhoZ+JRtUJElHzSqUnAtKk	dhion damanhuri	2026-06-13 16:14:49.659219	2026-06-13 16:19:22.368959
28	zovankacamata@gmail.com	zovan_zovankacamata	dataAnggota	$argon2id$v=19$m=65536,t=1,p=16$LE9XMehjT9R3c40ORKRJrA$Ob8hcGd1Y3ywoGqOEGkJ87tHNLrQEMpP20Az19psGrQ	zovan nya z	2026-06-14 03:54:37.412433	2026-06-14 03:54:37.412433
29	ipangacor@gmail.com	ipan_ipangacor	BPH	$argon2id$v=19$m=65536,t=1,p=16$u+Ryoz6JyT3AitjDFKVoEA$9HjNf9Ci8aX2SGckbi3Ajp/2TxHAShBiUMF0Q6Tym3w	ipan sang sdm	2026-06-15 05:10:54.151376	2026-06-15 05:12:29.482983
\.


--
-- Data for Name: work; Type: TABLE DATA; Schema: public; Owner: iko
--

COPY public.work (id, pengurus_id, title, tagline, description, slug, project_type, technologies, project_date, image_url, status, division, created_at, updated_at) FROM stdin;
2	2	Website Deteksi Femboy	Deteksi femboy tanpa ribet	bagi anda yang ingin mengetahui teman atau anda sendiri yang femboy, maka website ini cocok karena bisa deteksi femboy dengan mudah	website-femboy-deteksi	website	{html,css,javascript,mysql}	2026-01-17 00:00:00		published		2026-05-27 06:26:32.856417	2026-05-27 06:26:32.856418
3	3	Ricesource	Opensource ur ricing	platform to share ur linux configuration and using other people configuration on ur machine	wong sangar rak perlu rapat, seng penting garap jadi	web	{express.js,next.js,postgresql}	2024-12-20 00:00:00		draft		2026-05-27 06:26:32.857467	2026-05-27 06:26:32.857468
5	18	color pallete retro	retro style	kumpulan color pallete tema retro yang sudah siap untuk pakai dan pastinya terbaik dari yang terbaik, dan juga sudah sesuai dengan style jaman sekarang boss	color-palete-retro	website	{astro,figma,express.js}	2026-01-25 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-46a69f73-15b9-4fe5-b43c-7dff90243d07.jpg	draft	medcrev	2026-06-14 08:26:28.446392	2026-06-14 08:26:28.446392
7	1	sicon tools	tool pelacak orang	tools yang bisa digunakan untuk melacak orang dan mendapatkan data orang tersebut sehinga kamu bisa melakukan blackmail	spy-tools	script-tools	{python,rust,shell}	2026-01-25 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-929a5207-f8a7-4f4e-a977-7f1006eda381.jpg	published	bph	2026-06-14 09:21:50.089819	2026-06-14 09:21:50.08982
6	18	web pentester	pentester web learning	platform for you to learn how to pentest an a website	pentestter-web	website	{vue.js,rust,shell}	2026-01-25 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-21b7f939-bf37-4202-98ab-33b9fca0e0c7.jpg	unpublished	medcrev	2026-06-14 08:41:02.620407	2026-06-15 11:13:50.340217
8	1	d-form web	event management tool	tool untuk mengorganisasi event di sebuah organisasi atau instansi	event-tools-organization	form tools	{laravel,livewire,redis}	2026-01-25 00:00:00	http://localhost:9000/doscom-uploads/work/20260614-bf36757e-1b08-4a81-b3cd-4d4a5e0eac94.jpg	published	pemro	2026-06-14 12:01:27.888151	2026-06-15 10:53:35.793243
12	10	ricesource	sharing-platform	web untuk para pengguna linux bisa saling share ricing/custom linux mereka, dan user juga bisa pakai ricing/custom milik user lain secara gratis	linux-ricing-share	website	{next.js,postgresql,express.js}	2026-01-25 00:00:00	http://localhost:9000/doscom-uploads/work/20260615-0fa4e2fa-18e1-4951-bbf8-2ef7b5005066.jpg	rejected	pemro	2026-06-15 16:39:45.712074	2026-06-15 17:06:07.295224
4	\N	web pendeteksi femboy	femboy enjoyer	untuk mendeteksi apakah kamu itu seorang femboy atau tidak dengan menjawab beberapa pertanyaan dan juga dengan memberikan foto	web-femboy-enjoyer	website	{vue.js,postgresql,go}	2025-12-11 00:00:00	http://localhost:9000/doscom-uploads/gallery/20260613-1afcfa77-9c57-4bbe-b7b0-5b7b3c96ee86.jpg	draft	medcrev	2026-06-13 19:24:16.729621	2026-06-13 19:24:16.729621
\.


--
-- Data for Name: work_gallery; Type: TABLE DATA; Schema: public; Owner: iko
--

COPY public.work_gallery (id, id_work, id_gallery, created_at, updated_at) FROM stdin;
1	4	16	2026-06-13 19:24:16.733804	2026-06-13 19:24:16.733804
2	5	23	2026-06-14 08:26:28.450065	2026-06-14 08:26:28.450065
3	5	24	2026-06-14 08:26:28.450065	2026-06-14 08:26:28.450065
4	5	25	2026-06-14 08:26:28.450065	2026-06-14 08:26:28.450065
5	5	26	2026-06-14 08:26:28.450065	2026-06-14 08:26:28.450065
6	5	27	2026-06-14 08:26:28.450065	2026-06-14 08:26:28.450065
7	6	16	2026-06-14 08:41:02.620659	2026-06-14 08:41:02.620659
8	6	28	2026-06-14 08:41:02.620659	2026-06-14 08:41:02.620659
9	6	29	2026-06-14 08:41:02.620659	2026-06-14 08:41:02.620659
10	6	30	2026-06-14 08:41:02.620659	2026-06-14 08:41:02.620659
11	6	31	2026-06-14 08:41:02.620659	2026-06-14 08:41:02.620659
12	7	16	2026-06-14 09:21:50.091171	2026-06-14 09:21:50.091171
13	7	32	2026-06-14 09:21:50.091171	2026-06-14 09:21:50.091171
14	7	33	2026-06-14 09:21:50.091171	2026-06-14 09:21:50.091171
15	7	34	2026-06-14 09:21:50.091171	2026-06-14 09:21:50.091171
16	7	35	2026-06-14 09:21:50.091171	2026-06-14 09:21:50.091171
17	8	16	2026-06-14 12:01:27.890435	2026-06-14 12:01:27.890435
18	8	36	2026-06-14 12:01:27.890435	2026-06-14 12:01:27.890435
19	8	37	2026-06-14 12:01:27.890435	2026-06-14 12:01:27.890435
20	8	38	2026-06-14 12:01:27.890435	2026-06-14 12:01:27.890435
21	8	39	2026-06-14 12:01:27.890435	2026-06-14 12:01:27.890435
35	12	16	2026-06-15 16:39:45.713583	2026-06-15 16:39:45.713583
36	12	52	2026-06-15 16:39:45.713583	2026-06-15 16:39:45.713583
37	12	53	2026-06-15 16:39:45.713583	2026-06-15 16:39:45.713583
38	12	54	2026-06-15 16:39:45.713583	2026-06-15 16:39:45.713583
39	12	55	2026-06-15 16:39:45.713583	2026-06-15 16:39:45.713583
\.


--
-- Name: blog_gallery_id_seq; Type: SEQUENCE SET; Schema: public; Owner: iko
--

SELECT pg_catalog.setval('public.blog_gallery_id_seq', 1, false);


--
-- Name: blog_id_seq; Type: SEQUENCE SET; Schema: public; Owner: iko
--

SELECT pg_catalog.setval('public.blog_id_seq', 7, true);


--
-- Name: file_uploads_id_seq; Type: SEQUENCE SET; Schema: public; Owner: iko
--

SELECT pg_catalog.setval('public.file_uploads_id_seq', 73, true);


--
-- Name: gallery_id_seq; Type: SEQUENCE SET; Schema: public; Owner: iko
--

SELECT pg_catalog.setval('public.gallery_id_seq', 58, true);


--
-- Name: history_photos_id_seq; Type: SEQUENCE SET; Schema: public; Owner: iko
--

SELECT pg_catalog.setval('public.history_photos_id_seq', 1, false);


--
-- Name: history_timeline_id_seq; Type: SEQUENCE SET; Schema: public; Owner: iko
--

SELECT pg_catalog.setval('public.history_timeline_id_seq', 1, false);


--
-- Name: pengurus_id_seq; Type: SEQUENCE SET; Schema: public; Owner: iko
--

SELECT pg_catalog.setval('public.pengurus_id_seq', 20, true);


--
-- Name: pengurus_sosmed_id_seq; Type: SEQUENCE SET; Schema: public; Owner: iko
--

SELECT pg_catalog.setval('public.pengurus_sosmed_id_seq', 54, true);


--
-- Name: refresh_token_id_seq; Type: SEQUENCE SET; Schema: public; Owner: iko
--

SELECT pg_catalog.setval('public.refresh_token_id_seq', 89, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: iko
--

SELECT pg_catalog.setval('public.users_id_seq', 29, true);


--
-- Name: work_gallery_id_seq; Type: SEQUENCE SET; Schema: public; Owner: iko
--

SELECT pg_catalog.setval('public.work_gallery_id_seq', 39, true);


--
-- Name: work_id_seq; Type: SEQUENCE SET; Schema: public; Owner: iko
--

SELECT pg_catalog.setval('public.work_id_seq', 12, true);


--
-- Name: blog_gallery blog_gallery_pkey; Type: CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.blog_gallery
    ADD CONSTRAINT blog_gallery_pkey PRIMARY KEY (id);


--
-- Name: blog blog_pkey; Type: CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.blog
    ADD CONSTRAINT blog_pkey PRIMARY KEY (id);


--
-- Name: blog blog_slug_key; Type: CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.blog
    ADD CONSTRAINT blog_slug_key UNIQUE (slug);


--
-- Name: file_uploads file_uploads_pkey; Type: CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.file_uploads
    ADD CONSTRAINT file_uploads_pkey PRIMARY KEY (id);


--
-- Name: gallery gallery_pkey; Type: CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.gallery
    ADD CONSTRAINT gallery_pkey PRIMARY KEY (id);


--
-- Name: history_photos history_photos_pkey; Type: CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.history_photos
    ADD CONSTRAINT history_photos_pkey PRIMARY KEY (id);


--
-- Name: history_timeline history_timeline_pkey; Type: CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.history_timeline
    ADD CONSTRAINT history_timeline_pkey PRIMARY KEY (id);


--
-- Name: pengurus pengurus_email_key; Type: CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.pengurus
    ADD CONSTRAINT pengurus_email_key UNIQUE (email);


--
-- Name: pengurus pengurus_pkey; Type: CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.pengurus
    ADD CONSTRAINT pengurus_pkey PRIMARY KEY (id);


--
-- Name: pengurus_sosmed pengurus_sosmed_pkey; Type: CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.pengurus_sosmed
    ADD CONSTRAINT pengurus_sosmed_pkey PRIMARY KEY (id);


--
-- Name: refresh_token refresh_token_pkey; Type: CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.refresh_token
    ADD CONSTRAINT refresh_token_pkey PRIMARY KEY (id);


--
-- Name: refresh_token refresh_token_user_id_key; Type: CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.refresh_token
    ADD CONSTRAINT refresh_token_user_id_key UNIQUE (user_id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: work_gallery work_gallery_pkey; Type: CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.work_gallery
    ADD CONSTRAINT work_gallery_pkey PRIMARY KEY (id);


--
-- Name: work work_pkey; Type: CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.work
    ADD CONSTRAINT work_pkey PRIMARY KEY (id);


--
-- Name: idx_blog_thumbnail_url; Type: INDEX; Schema: public; Owner: iko
--

CREATE INDEX idx_blog_thumbnail_url ON public.blog USING btree (thumbnail_url);


--
-- Name: idx_file_uploads_category; Type: INDEX; Schema: public; Owner: iko
--

CREATE INDEX idx_file_uploads_category ON public.file_uploads USING btree (category);


--
-- Name: idx_file_uploads_uploaded_at; Type: INDEX; Schema: public; Owner: iko
--

CREATE INDEX idx_file_uploads_uploaded_at ON public.file_uploads USING btree (uploaded_at DESC);


--
-- Name: idx_file_uploads_user_id; Type: INDEX; Schema: public; Owner: iko
--

CREATE INDEX idx_file_uploads_user_id ON public.file_uploads USING btree (user_id);


--
-- Name: idx_history_photos_history_id; Type: INDEX; Schema: public; Owner: iko
--

CREATE INDEX idx_history_photos_history_id ON public.history_photos USING btree (id_history);


--
-- Name: idx_history_timeline_disyplay_order; Type: INDEX; Schema: public; Owner: iko
--

CREATE INDEX idx_history_timeline_disyplay_order ON public.history_timeline USING btree (display_order);


--
-- Name: idx_pengurus_photo_url; Type: INDEX; Schema: public; Owner: iko
--

CREATE INDEX idx_pengurus_photo_url ON public.pengurus USING btree (photo_url);


--
-- Name: idx_refresh_tn_user_id; Type: INDEX; Schema: public; Owner: iko
--

CREATE INDEX idx_refresh_tn_user_id ON public.refresh_token USING btree (user_id);


--
-- Name: idx_work_status; Type: INDEX; Schema: public; Owner: iko
--

CREATE INDEX idx_work_status ON public.work USING btree (status);


--
-- Name: idx_work_technologies_gin; Type: INDEX; Schema: public; Owner: iko
--

CREATE INDEX idx_work_technologies_gin ON public.work USING gin (technologies);


--
-- Name: blog blog_author_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.blog
    ADD CONSTRAINT blog_author_id_fkey FOREIGN KEY (author_id) REFERENCES public.pengurus(id);


--
-- Name: blog_gallery blog_gallery_id_blog_fkey; Type: FK CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.blog_gallery
    ADD CONSTRAINT blog_gallery_id_blog_fkey FOREIGN KEY (id_blog) REFERENCES public.blog(id);


--
-- Name: blog_gallery blog_gallery_id_gallery_fkey; Type: FK CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.blog_gallery
    ADD CONSTRAINT blog_gallery_id_gallery_fkey FOREIGN KEY (id_gallery) REFERENCES public.gallery(id);


--
-- Name: file_uploads file_uploads_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.file_uploads
    ADD CONSTRAINT file_uploads_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: gallery gallery_file_upload_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.gallery
    ADD CONSTRAINT gallery_file_upload_id_fkey FOREIGN KEY (file_upload_id) REFERENCES public.file_uploads(id);


--
-- Name: gallery gallery_id_users_fkey; Type: FK CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.gallery
    ADD CONSTRAINT gallery_id_users_fkey FOREIGN KEY (id_users) REFERENCES public.users(id) ON DELETE SET NULL;


--
-- Name: history_photos history_photos_id_history_fkey; Type: FK CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.history_photos
    ADD CONSTRAINT history_photos_id_history_fkey FOREIGN KEY (id_history) REFERENCES public.history_timeline(id);


--
-- Name: history_timeline history_timeline_id_author_fkey; Type: FK CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.history_timeline
    ADD CONSTRAINT history_timeline_id_author_fkey FOREIGN KEY (id_author) REFERENCES public.users(id);


--
-- Name: pengurus pengurus_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.pengurus
    ADD CONSTRAINT pengurus_id_user_fkey FOREIGN KEY (id_user) REFERENCES public.users(id) ON DELETE SET NULL;


--
-- Name: pengurus_sosmed pengurus_sosmed_pengurus_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.pengurus_sosmed
    ADD CONSTRAINT pengurus_sosmed_pengurus_id_fkey FOREIGN KEY (pengurus_id) REFERENCES public.pengurus(id) ON DELETE CASCADE;


--
-- Name: work_gallery work_gallery_id_gallery_fkey; Type: FK CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.work_gallery
    ADD CONSTRAINT work_gallery_id_gallery_fkey FOREIGN KEY (id_gallery) REFERENCES public.gallery(id);


--
-- Name: work_gallery work_gallery_id_work_fkey; Type: FK CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.work_gallery
    ADD CONSTRAINT work_gallery_id_work_fkey FOREIGN KEY (id_work) REFERENCES public.work(id);


--
-- Name: work work_pengurus_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.work
    ADD CONSTRAINT work_pengurus_id_fkey FOREIGN KEY (pengurus_id) REFERENCES public.pengurus(id) ON DELETE SET NULL;


--
-- PostgreSQL database dump complete
--

\unrestrict JHgEa2pXsuyJWQOQc7vstfhyb1phOEWgOUpTett2M0cNLpegp6s7B23RNHuIdrz

