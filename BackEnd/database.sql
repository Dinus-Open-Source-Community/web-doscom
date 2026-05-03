--
-- PostgreSQL database dump
--

\restrict qYNhcBuUUeVeUyCNlCSNBcrT5MYLHBYKrxEtPxkgJJ6vKuPfdtePztVg0z36A7n

-- Dumped from database version 18.3
-- Dumped by pg_dump version 18.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
    'scheduled'
);


ALTER TYPE public.blog_status OWNER TO iko;

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
    id_gallery integer
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
    sosmed character varying(50),
    period character varying(50),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pengurus_divisi_check CHECK (((divisi)::text = ANY ((ARRAY['bph'::character varying, 'pemro'::character varying, 'jaringan'::character varying, 'medcrev'::character varying, 'data'::character varying])::text[]))),
    CONSTRAINT pengurus_position_check CHECK ((("position")::text = ANY ((ARRAY['ketum'::character varying, 'sdm'::character varying, 'pr'::character varying, 'pm'::character varying, 'pmAng'::character varying, 'sekum'::character varying, 'bendum'::character varying, 'sekAng'::character varying, 'bendAng'::character varying, 'KoorPemro'::character varying, 'KoorJaringan'::character varying, 'KoorMedcrev'::character varying, 'KoorData'::character varying, 'anggotaAktif'::character varying, 'PemroAng'::character varying, 'JaringanAng'::character varying, 'MedcrevAng'::character varying, 'DataAng'::character varying])::text[]))),
    CONSTRAINT pengurus_sosmed_check CHECK (((sosmed)::text = ANY ((ARRAY['instagram'::character varying, 'linkedin'::character varying, 'github'::character varying])::text[])))
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
-- Name: refresh_token; Type: TABLE; Schema: public; Owner: iko
--

CREATE TABLE public.refresh_token (
    id integer NOT NULL,
    user_id integer NOT NULL,
    token_hash character varying(255) NOT NULL,
    expires_at timestamp without time zone NOT NULL,
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
    image_url text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.work OWNER TO iko;

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
-- Data for Name: blog; Type: TABLE DATA; Schema: public; Owner: iko
--

COPY public.blog (id, author_id, title, slug, content, thumbnail_url, kategori, published_at, status, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: blog_gallery; Type: TABLE DATA; Schema: public; Owner: iko
--

COPY public.blog_gallery (id, id_blog, id_gallery) FROM stdin;
\.


--
-- Data for Name: file_uploads; Type: TABLE DATA; Schema: public; Owner: iko
--

COPY public.file_uploads (id, user_id, category, original_filename, stored_filename, file_size, content_type, file_url, uploaded_at, updated_at) FROM stdin;
\.


--
-- Data for Name: gallery; Type: TABLE DATA; Schema: public; Owner: iko
--

COPY public.gallery (id, id_users, file_upload_id, gallery_name, gallery_type, description, event_date, file_url, created_at, updated_at) FROM stdin;
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

COPY public.pengurus (id, id_user, photo_url, email, divisi, name, "position", sosmed, period, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: refresh_token; Type: TABLE DATA; Schema: public; Owner: iko
--

COPY public.refresh_token (id, user_id, token_hash, expires_at, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: iko
--

COPY public.schema_migrations (version, dirty) FROM stdin;
10	f
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: iko
--

COPY public.users (id, email, username, role, password, full_name, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: work; Type: TABLE DATA; Schema: public; Owner: iko
--

COPY public.work (id, pengurus_id, title, tagline, description, slug, project_type, technologies, image_url, created_at, updated_at) FROM stdin;
\.


--
-- Name: blog_gallery_id_seq; Type: SEQUENCE SET; Schema: public; Owner: iko
--

SELECT pg_catalog.setval('public.blog_gallery_id_seq', 1, false);


--
-- Name: blog_id_seq; Type: SEQUENCE SET; Schema: public; Owner: iko
--

SELECT pg_catalog.setval('public.blog_id_seq', 1, false);


--
-- Name: file_uploads_id_seq; Type: SEQUENCE SET; Schema: public; Owner: iko
--

SELECT pg_catalog.setval('public.file_uploads_id_seq', 1, false);


--
-- Name: gallery_id_seq; Type: SEQUENCE SET; Schema: public; Owner: iko
--

SELECT pg_catalog.setval('public.gallery_id_seq', 1, false);


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

SELECT pg_catalog.setval('public.pengurus_id_seq', 1, false);


--
-- Name: refresh_token_id_seq; Type: SEQUENCE SET; Schema: public; Owner: iko
--

SELECT pg_catalog.setval('public.refresh_token_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: iko
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- Name: work_id_seq; Type: SEQUENCE SET; Schema: public; Owner: iko
--

SELECT pg_catalog.setval('public.work_id_seq', 1, false);


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
-- Name: idx_work_image_url; Type: INDEX; Schema: public; Owner: iko
--

CREATE INDEX idx_work_image_url ON public.work USING btree (image_url);


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
    ADD CONSTRAINT gallery_id_users_fkey FOREIGN KEY (id_users) REFERENCES public.users(id);


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
    ADD CONSTRAINT pengurus_id_user_fkey FOREIGN KEY (id_user) REFERENCES public.users(id);


--
-- Name: work work_pengurus_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: iko
--

ALTER TABLE ONLY public.work
    ADD CONSTRAINT work_pengurus_id_fkey FOREIGN KEY (pengurus_id) REFERENCES public.pengurus(id);


--
-- PostgreSQL database dump complete
--

\unrestrict qYNhcBuUUeVeUyCNlCSNBcrT5MYLHBYKrxEtPxkgJJ6vKuPfdtePztVg0z36A7n

