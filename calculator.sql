PGDMP     9    .                z        
   calculator    14.4    14.4 	    �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16418 
   calculator    DATABASE     n   CREATE DATABASE calculator WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'English_United States.1252';
    DROP DATABASE calculator;
                postgres    false            �            1259    16420    records    TABLE     �   CREATE TABLE public.records (
    id integer NOT NULL,
    first_number double precision,
    second_number double precision,
    operator character(1),
    answer double precision
);
    DROP TABLE public.records;
       public         heap    postgres    false            �            1259    16419    records_id_seq    SEQUENCE     �   ALTER TABLE public.records ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.records_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    210            �          0    16420    records 
   TABLE DATA           T   COPY public.records (id, first_number, second_number, operator, answer) FROM stdin;
    public          postgres    false    210   �       �           0    0    records_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.records_id_seq', 151, true);
          public          postgres    false    209            ]           2606    16426    records records_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public.records
    ADD CONSTRAINT records_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY public.records DROP CONSTRAINT records_pkey;
       public            postgres    false    210            �   N   x�̻�0��T
��N?�^�<N6x�2'�x`���(_I����04����va��8�Fen&b_����%����     