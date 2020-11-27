--
-- PostgreSQL database dump
--

-- Dumped from database version 12.5 (Ubuntu 12.5-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 12.5 (Ubuntu 12.5-0ubuntu0.20.04.1)

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
-- Name: acc_status; Type: TYPE; Schema: public; Owner: bsabre
--

CREATE TYPE public.acc_status AS ENUM (
    'confirmed',
    'not confirmed',
    ''
);


ALTER TYPE public.acc_status OWNER TO bsabre;

--
-- Name: enum_gender; Type: TYPE; Schema: public; Owner: bsabre
--

CREATE TYPE public.enum_gender AS ENUM (
    'male',
    'female',
    ''
);


ALTER TYPE public.enum_gender OWNER TO bsabre;

--
-- Name: enum_orientation; Type: TYPE; Schema: public; Owner: bsabre
--

CREATE TYPE public.enum_orientation AS ENUM (
    'hetero',
    'bi',
    'homo',
    ''
);


ALTER TYPE public.enum_orientation OWNER TO bsabre;

--
-- Name: enum_status; Type: TYPE; Schema: public; Owner: bsabre
--

CREATE TYPE public.enum_status AS ENUM (
    'confirmed',
    'not confirmed',
    ''
);


ALTER TYPE public.enum_status OWNER TO bsabre;

--
-- Name: matcha_gender; Type: TYPE; Schema: public; Owner: bsabre
--

CREATE TYPE public.matcha_gender AS ENUM (
    'male',
    'female',
    ''
);


ALTER TYPE public.matcha_gender OWNER TO bsabre;

--
-- Name: matcha_orientation; Type: TYPE; Schema: public; Owner: bsabre
--

CREATE TYPE public.matcha_orientation AS ENUM (
    'getero',
    'bi',
    'gay',
    ''
);


ALTER TYPE public.matcha_orientation OWNER TO bsabre;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: claims; Type: TABLE; Schema: public; Owner: bsabre
--

CREATE TABLE public.claims (
    uidsender integer NOT NULL,
    uidreceiver integer NOT NULL
);


ALTER TABLE public.claims OWNER TO bsabre;

--
-- Name: devices; Type: TABLE; Schema: public; Owner: bsabre
--

CREATE TABLE public.devices (
    id integer NOT NULL,
    uid integer NOT NULL,
    device character varying(150) NOT NULL
);


ALTER TABLE public.devices OWNER TO bsabre;

--
-- Name: devices_id_seq; Type: SEQUENCE; Schema: public; Owner: bsabre
--

CREATE SEQUENCE public.devices_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.devices_id_seq OWNER TO bsabre;

--
-- Name: devices_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: bsabre
--

ALTER SEQUENCE public.devices_id_seq OWNED BY public.devices.id;


--
-- Name: history; Type: TABLE; Schema: public; Owner: bsabre
--

CREATE TABLE public.history (
    id integer NOT NULL,
    uid integer NOT NULL,
    targetuid integer NOT NULL,
    "time" timestamp without time zone NOT NULL
);


ALTER TABLE public.history OWNER TO bsabre;

--
-- Name: history_id_seq; Type: SEQUENCE; Schema: public; Owner: bsabre
--

CREATE SEQUENCE public.history_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.history_id_seq OWNER TO bsabre;

--
-- Name: history_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: bsabre
--

ALTER SEQUENCE public.history_id_seq OWNED BY public.history.id;


--
-- Name: ignores; Type: TABLE; Schema: public; Owner: bsabre
--

CREATE TABLE public.ignores (
    uidsender integer NOT NULL,
    uidreceiver integer NOT NULL
);


ALTER TABLE public.ignores OWNER TO bsabre;

--
-- Name: interests; Type: TABLE; Schema: public; Owner: bsabre
--

CREATE TABLE public.interests (
    id integer NOT NULL,
    name character varying(100) NOT NULL
);


ALTER TABLE public.interests OWNER TO bsabre;

--
-- Name: interests_id_seq; Type: SEQUENCE; Schema: public; Owner: bsabre
--

CREATE SEQUENCE public.interests_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.interests_id_seq OWNER TO bsabre;

--
-- Name: interests_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: bsabre
--

ALTER SEQUENCE public.interests_id_seq OWNED BY public.interests.id;


--
-- Name: likes; Type: TABLE; Schema: public; Owner: bsabre
--

CREATE TABLE public.likes (
    uidsender integer NOT NULL,
    uidreceiver integer NOT NULL
);


ALTER TABLE public.likes OWNER TO bsabre;

--
-- Name: messages; Type: TABLE; Schema: public; Owner: bsabre
--

CREATE TABLE public.messages (
    mid integer NOT NULL,
    uidsender integer NOT NULL,
    uidreceiver integer NOT NULL,
    body character varying(300) NOT NULL,
    active boolean DEFAULT true
);


ALTER TABLE public.messages OWNER TO bsabre;

--
-- Name: messages_mid_seq; Type: SEQUENCE; Schema: public; Owner: bsabre
--

CREATE SEQUENCE public.messages_mid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.messages_mid_seq OWNER TO bsabre;

--
-- Name: messages_mid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: bsabre
--

ALTER SEQUENCE public.messages_mid_seq OWNED BY public.messages.mid;


--
-- Name: notifs; Type: TABLE; Schema: public; Owner: bsabre
--

CREATE TABLE public.notifs (
    nid integer NOT NULL,
    uidsender integer NOT NULL,
    uidreceiver integer NOT NULL,
    body character varying(250) NOT NULL
);


ALTER TABLE public.notifs OWNER TO bsabre;

--
-- Name: notifs_nid_seq; Type: SEQUENCE; Schema: public; Owner: bsabre
--

CREATE SEQUENCE public.notifs_nid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.notifs_nid_seq OWNER TO bsabre;

--
-- Name: notifs_nid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: bsabre
--

ALTER SEQUENCE public.notifs_nid_seq OWNED BY public.notifs.nid;


--
-- Name: photos; Type: TABLE; Schema: public; Owner: bsabre
--

CREATE TABLE public.photos (
    pid integer NOT NULL,
    uid integer NOT NULL,
    src character varying(10000000) NOT NULL
);


ALTER TABLE public.photos OWNER TO bsabre;

--
-- Name: photos_pid_seq; Type: SEQUENCE; Schema: public; Owner: bsabre
--

CREATE SEQUENCE public.photos_pid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.photos_pid_seq OWNER TO bsabre;

--
-- Name: photos_pid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: bsabre
--

ALTER SEQUENCE public.photos_pid_seq OWNED BY public.photos.pid;


--
-- Name: users; Type: TABLE; Schema: public; Owner: bsabre
--

CREATE TABLE public.users (
    uid integer NOT NULL,
    mail character varying(30) DEFAULT ''::character varying NOT NULL,
    encryptedpass character varying(35) NOT NULL,
    fname character varying(30) DEFAULT ''::character varying NOT NULL,
    lname character varying(30) DEFAULT ''::character varying NOT NULL,
    birth date,
    gender public.enum_gender DEFAULT ''::public.enum_gender NOT NULL,
    orientation public.enum_orientation DEFAULT ''::public.enum_orientation NOT NULL,
    bio character varying(300) DEFAULT ''::character varying NOT NULL,
    avaid integer,
    latitude double precision,
    longitude double precision,
    interests character varying(100)[] DEFAULT '{}'::character varying[],
    status public.enum_status DEFAULT 'not confirmed'::public.enum_status NOT NULL,
    search_visibility boolean DEFAULT false NOT NULL,
    rating integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.users OWNER TO bsabre;

--
-- Name: users_uid_seq; Type: SEQUENCE; Schema: public; Owner: bsabre
--

CREATE SEQUENCE public.users_uid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_uid_seq OWNER TO bsabre;

--
-- Name: users_uid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: bsabre
--

ALTER SEQUENCE public.users_uid_seq OWNED BY public.users.uid;


--
-- Name: devices id; Type: DEFAULT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.devices ALTER COLUMN id SET DEFAULT nextval('public.devices_id_seq'::regclass);


--
-- Name: history id; Type: DEFAULT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.history ALTER COLUMN id SET DEFAULT nextval('public.history_id_seq'::regclass);


--
-- Name: interests id; Type: DEFAULT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.interests ALTER COLUMN id SET DEFAULT nextval('public.interests_id_seq'::regclass);


--
-- Name: messages mid; Type: DEFAULT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.messages ALTER COLUMN mid SET DEFAULT nextval('public.messages_mid_seq'::regclass);


--
-- Name: notifs nid; Type: DEFAULT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.notifs ALTER COLUMN nid SET DEFAULT nextval('public.notifs_nid_seq'::regclass);


--
-- Name: photos pid; Type: DEFAULT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.photos ALTER COLUMN pid SET DEFAULT nextval('public.photos_pid_seq'::regclass);


--
-- Name: users uid; Type: DEFAULT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.users ALTER COLUMN uid SET DEFAULT nextval('public.users_uid_seq'::regclass);


--
-- Data for Name: claims; Type: TABLE DATA; Schema: public; Owner: bsabre
--

COPY public.claims (uidsender, uidreceiver) FROM stdin;
\.


--
-- Data for Name: devices; Type: TABLE DATA; Schema: public; Owner: bsabre
--

COPY public.devices (id, uid, device) FROM stdin;
\.


--
-- Data for Name: history; Type: TABLE DATA; Schema: public; Owner: bsabre
--

COPY public.history (id, uid, targetuid, "time") FROM stdin;
\.


--
-- Data for Name: ignores; Type: TABLE DATA; Schema: public; Owner: bsabre
--

COPY public.ignores (uidsender, uidreceiver) FROM stdin;
1	1
2	2
3	3
4	4
5	5
6	6
7	7
8	8
9	9
10	10
11	11
12	12
13	13
14	14
15	15
16	16
17	17
18	18
19	19
20	20
21	21
22	22
23	23
24	24
25	25
26	26
27	27
28	28
29	29
30	30
31	31
32	32
33	33
34	34
35	35
36	36
37	37
38	38
39	39
40	40
41	41
42	42
43	43
44	44
45	45
46	46
47	47
48	48
49	49
50	50
51	51
52	52
53	53
54	54
55	55
56	56
57	57
58	58
59	59
60	60
61	61
62	62
63	63
64	64
65	65
66	66
67	67
68	68
69	69
70	70
71	71
72	72
73	73
74	74
75	75
76	76
77	77
78	78
79	79
80	80
81	81
82	82
83	83
84	84
85	85
86	86
87	87
88	88
89	89
90	90
91	91
92	92
93	93
94	94
95	95
96	96
97	97
98	98
99	99
100	100
101	101
102	102
103	103
104	104
105	105
106	106
107	107
108	108
109	109
110	110
111	111
112	112
113	113
114	114
115	115
116	116
117	117
118	118
119	119
120	120
121	121
122	122
123	123
124	124
125	125
126	126
127	127
128	128
129	129
130	130
131	131
132	132
133	133
134	134
135	135
136	136
137	137
138	138
139	139
140	140
141	141
142	142
143	143
144	144
145	145
146	146
147	147
148	148
149	149
150	150
151	151
152	152
153	153
154	154
155	155
156	156
157	157
158	158
159	159
160	160
161	161
162	162
163	163
164	164
165	165
166	166
167	167
168	168
169	169
170	170
171	171
172	172
173	173
174	174
175	175
176	176
177	177
178	178
179	179
180	180
181	181
182	182
183	183
184	184
185	185
186	186
187	187
188	188
189	189
190	190
191	191
192	192
193	193
194	194
195	195
196	196
197	197
198	198
199	199
200	200
201	201
202	202
203	203
204	204
205	205
206	206
207	207
208	208
209	209
210	210
211	211
212	212
213	213
214	214
215	215
216	216
217	217
218	218
219	219
220	220
221	221
222	222
223	223
224	224
225	225
226	226
227	227
228	228
229	229
230	230
231	231
232	232
233	233
234	234
235	235
236	236
237	237
238	238
239	239
240	240
241	241
242	242
243	243
244	244
245	245
246	246
247	247
248	248
249	249
250	250
251	251
252	252
253	253
254	254
255	255
256	256
257	257
258	258
259	259
260	260
261	261
262	262
263	263
264	264
265	265
266	266
267	267
268	268
269	269
270	270
271	271
272	272
273	273
274	274
275	275
276	276
277	277
278	278
279	279
280	280
281	281
282	282
283	283
284	284
285	285
286	286
287	287
288	288
289	289
290	290
291	291
292	292
293	293
294	294
295	295
296	296
297	297
298	298
299	299
300	300
301	301
302	302
303	303
304	304
305	305
306	306
307	307
308	308
309	309
310	310
311	311
312	312
313	313
314	314
315	315
316	316
317	317
318	318
319	319
320	320
321	321
322	322
323	323
324	324
325	325
326	326
327	327
328	328
329	329
330	330
331	331
332	332
333	333
334	334
335	335
336	336
337	337
338	338
339	339
340	340
341	341
342	342
343	343
344	344
345	345
346	346
347	347
348	348
349	349
350	350
351	351
352	352
353	353
354	354
355	355
356	356
357	357
358	358
359	359
360	360
361	361
362	362
363	363
364	364
365	365
366	366
367	367
368	368
369	369
370	370
371	371
372	372
373	373
374	374
375	375
376	376
377	377
378	378
379	379
380	380
381	381
382	382
383	383
384	384
385	385
386	386
387	387
388	388
389	389
390	390
391	391
392	392
393	393
394	394
395	395
396	396
397	397
398	398
399	399
400	400
401	401
402	402
403	403
404	404
405	405
406	406
407	407
408	408
409	409
410	410
411	411
412	412
413	413
414	414
415	415
416	416
417	417
418	418
419	419
420	420
421	421
422	422
423	423
424	424
425	425
426	426
427	427
428	428
429	429
430	430
431	431
432	432
433	433
434	434
435	435
436	436
437	437
438	438
439	439
440	440
441	441
442	442
443	443
444	444
445	445
446	446
447	447
448	448
449	449
450	450
451	451
452	452
453	453
454	454
455	455
456	456
457	457
458	458
459	459
460	460
461	461
462	462
463	463
464	464
465	465
466	466
467	467
468	468
469	469
470	470
471	471
472	472
473	473
474	474
475	475
476	476
477	477
478	478
479	479
480	480
481	481
482	482
483	483
484	484
485	485
486	486
487	487
488	488
489	489
490	490
491	491
492	492
493	493
494	494
495	495
496	496
497	497
498	498
499	499
500	500
501	501
\.


--
-- Data for Name: interests; Type: TABLE DATA; Schema: public; Owner: bsabre
--

COPY public.interests (id, name) FROM stdin;
1	culture
2	programming
3	find something new
4	youtube
5	drink beer
6	architecture
7	cooking
8	video games
9	politics
10	football
\.


--
-- Data for Name: likes; Type: TABLE DATA; Schema: public; Owner: bsabre
--

COPY public.likes (uidsender, uidreceiver) FROM stdin;
\.


--
-- Data for Name: messages; Type: TABLE DATA; Schema: public; Owner: bsabre
--

COPY public.messages (mid, uidsender, uidreceiver, body, active) FROM stdin;
\.


--
-- Data for Name: notifs; Type: TABLE DATA; Schema: public; Owner: bsabre
--

COPY public.notifs (nid, uidsender, uidreceiver, body) FROM stdin;
\.


--
-- Data for Name: photos; Type: TABLE DATA; Schema: public; Owner: bsabre
--

COPY public.photos (pid, uid, src) FROM stdin;
1	2	https://sun1-87.userapi.com/impf/c623322/v623322000/1adde/HOu2FKMgK5g.jpg?size=200x0&quality=90&crop=193,172,1063,1086&sign=db1a98d9231f2e2887662535479bd00e&ava=1
2	3	https://sun1-17.userapi.com/impf/34UEdlxioTPJKUTsWz9CcsRPxYi6vnGqg7_kEw/uL0zBsevqEI.jpg?size=200x0&quality=90&crop=0,0,1024,654&rotate=270&sign=9f2b41e263d7934005c78315f46e02ad&ava=1
3	4	https://sun1-98.userapi.com/impf/c849136/v849136110/dea70/iWTZgsURaBU.jpg?size=200x0&quality=90&crop=76,37,897,897&sign=6e9aef3cc0eb686066cc2d5e9c6fad49&ava=1
4	5	https://sun9-69.userapi.com/c9625/u555015/a_cbc8e53e.jpg?ava=1
5	6	https://sun1-94.userapi.com/impg/4WLi5wDJkXA7N357ROpQkpzoYHhMH4Rvz3hKjw/28ph6FRiOf0.jpg?size=200x0&quality=90&crop=1,126,999,999&sign=d587707e87aabc65c9daaa446c149d94&ava=1
6	7	https://sun1-27.userapi.com/impf/c836532/v836532023/2d9d/e9nrUTCQAeA.jpg?size=200x0&quality=90&crop=71,71,1523,2017&sign=bc979fb35d2d717e74a571396ee80481&ava=1
7	8	https://sun1-15.userapi.com/impf/SGWNTjatoRxbIVCGauE6J63Vm_X12vqjFAR6iA/YsL_x4gxv_w.jpg?size=200x0&quality=90&crop=0,0,640,640&sign=9b43eeff555f5a6935b72dddf1fbe8aa&ava=1
8	9	https://sun1-47.userapi.com/impf/c604828/v604828032/1f0e1/r52Gz9JeHzI.jpg?size=200x0&quality=90&crop=32,55,626,937&sign=35b5bd731ec25cd809531249a2b3fe82&ava=1
9	10	https://sun9-21.userapi.com/c30/u555034/a_68db014.jpg?ava=1
10	11	https://sun9-7.userapi.com/c10/u555038/a_d9a5308.jpg?ava=1
11	12	https://sun1-16.userapi.com/impf/TPk5qp0zkOzHMFZIoCkXymS1hJ7r12Hh1fCH4Q/i1apnKzZ3Bw.jpg?size=200x0&quality=90&crop=0,0,480,640&sign=133da0aa55745f69dcb18db6390cb570&ava=1
12	13	https://sun1-93.userapi.com/impf/c850336/v850336780/4cf37/N1y1pdrlpnY.jpg?size=200x0&quality=90&crop=0,0,1532,1532&sign=a2ffb926826c43a1d48cecd049a6e862&ava=1
13	14	https://sun1-96.userapi.com/impg/c857732/v857732090/1f274d/1DZxV3Ll4Tw.jpg?size=200x0&quality=90&crop=0,269,540,540&sign=80362648bd851afd5dec1249ace38483&ava=1
14	15	https://sun1-96.userapi.com/impf/c629504/v629504047/8136/HeNDx2rrUOE.jpg?size=200x0&quality=90&crop=67,67,1017,1913&sign=7cee66e0e39265c469e524b24c8f4d2d&ava=1
15	16	https://sun1-96.userapi.com/impf/m5IFGuwMJyz8gDmtimD6Q9uJiNWBaa543C7knw/twOw8d7VFFA.jpg?size=200x0&quality=90&crop=67,67,1391,1913&sign=2fe5dd13db336f75134ebc14840f1197&ava=1
16	17	https://sun1-98.userapi.com/impg/nP76XKpHGP9hsTb_7pHYyEZrgb98tqelGRKbRA/0OINZV1SXgQ.jpg?size=200x0&quality=90&crop=97,0,401,401&sign=0af661917ada3d5f3b7088a728f98054&ava=1
17	18	https://sun1-17.userapi.com/impf/wLuWDHjG7vSKIZArZs4nuqUjdMZ8Lpp5mx0bQw/R7pw9vQc7c4.jpg?size=200x0&quality=90&crop=224,84,876,876&sign=8acee313b8c0724f316d43aa570344cf&ava=1
18	19	https://sun9-20.userapi.com/c11/u555072/a_53453f5.jpg?ava=1
19	20	https://sun1-16.userapi.com/impg/R_pJ2P8OUC0yQMYxZGEOMHgx0XlwUY8r8k9eZg/jqlzBeA3R3M.jpg?size=200x0&quality=90&crop=161,67,1385,1385&sign=c12c6c92f33d2523026e7a581aebe81b&ava=1
20	21	https://sun1-93.userapi.com/impf/c627918/v627918079/18d47/upyoYOPa4KM.jpg?size=200x0&quality=90&crop=0,0,959,960&sign=ed6818133f58c015b2510b4ae9f3075c&ava=1
21	22	https://sun1-93.userapi.com/impf/c840330/v840330245/7e508/JrvaVlq2c0g.jpg?size=200x0&quality=90&crop=18,89,364,364&sign=0d47f83b78ff3a97414a6119e11b2861&ava=1
22	23	https://sun1-87.userapi.com/impf/c830400/v830400630/1cadb0/n8pEnwMRFLo.jpg?size=200x0&quality=90&crop=0,319,960,960&sign=06196d35b2ebb8c95ca83a96a68cb841&ava=1
23	24	https://sun1-29.userapi.com/impf/c836433/v836433089/53c72/eyQ2-AIuaaY.jpg?size=200x0&quality=90&crop=0,0,667,1000&sign=7983a6825ed3c628d7b02dfd3555c93c&ava=1
24	25	https://sun1-93.userapi.com/impg/c858328/v858328272/151034/7Ew_S-NdbXg.jpg?size=200x0&quality=90&crop=0,142,1620,1620&sign=7092d1c0077994735e0f017a51479d7a&ava=1
25	26	https://sun1-29.userapi.com/impf/c853424/v853424565/cebd9/9zPFmxIvGJ4.jpg?size=200x0&quality=90&crop=0,202,1218,1219&sign=4266533b2cd6519ab1d3b7ec3dcef563&ava=1
26	27	https://sun9-43.userapi.com/c10356/u555099/a_eaf4d77d.jpg?ava=1
27	28	https://sun9-20.userapi.com/c14/u555100/a_3e82f41.jpg?ava=1
28	29	https://sun9-32.userapi.com/c11381/u555102/a_5d658abc.jpg?ava=1
29	30	https://sun1-85.userapi.com/impf/E-8y9L-XVuh464RJTpWKblKVOQXqvy1fb9mM7g/MSkO1xRyz_I.jpg?size=200x0&quality=90&crop=398,0,876,876&sign=f42cabebad84f2e49d449fb97da3e3a1&ava=1
30	31	https://sun1-95.userapi.com/impf/c841539/v841539498/37072/KtnaRl-AkcY.jpg?size=200x0&quality=90&crop=342,596,1023,1023&sign=87db18f0ed6e072aba53ac482f5aaf75&ava=1
31	32	https://sun9-13.userapi.com/c18/u555121/a_d7be21a.jpg?ava=1
32	33	https://sun1-90.userapi.com/impf/c623721/v623721124/2fd84/5x0wSMzHi9g.jpg?size=200x0&quality=90&crop=122,8,1699,1699&sign=2dad499b0b1368626636cab3f1f9acac&ava=1
33	34	https://sun9-15.userapi.com/c10/u555126/a_3ee9f30.jpg?ava=1
34	35	https://sun1-26.userapi.com/impf/c840521/v840521716/1c576/Jypqvpbotvc.jpg?size=200x0&quality=90&crop=0,117,960,960&sign=c39066ebbbb3bc08eba69a38b6c5f397&ava=1
35	36	https://sun1-30.userapi.com/impf/A_YvJw3GZdsje0xYKmBOkY8QJXQQ-UlJEbb7Bw/lB5d-lgtdLg.jpg?size=200x0&quality=90&crop=0,20,200,227&sign=424d3bafd4f1f31557ac5d6c55b8b193&ava=1
36	37	https://sun1-30.userapi.com/GIChmJV0EpnxEIocBWkn1fgsjUESWozWdlrc4A/C2dLn7yzcg0.jpg?ava=1
37	38	https://sun1-96.userapi.com/impf/c847220/v847220399/4418a/gNSKIUGf07w.jpg?size=200x0&quality=90&crop=0,0,983,983&sign=ef77d0944f54d4c2710b25b4088f8e5c&ava=1
38	39	https://sun1-18.userapi.com/impf/IBL_hG4hs2qyNXUHjFX2DDQjjW9kf7v5O-Gx5A/EdWnmoQOH7I.jpg?size=200x0&quality=90&crop=0,0,370,445&sign=e1655ed165381f3cc1869ff436a40653&ava=1
39	40	https://sun9-67.userapi.com/c149/u555143/a_81516c15.jpg?ava=1
40	41	https://sun1-20.userapi.com/impf/vvgyxxtcuKJHliDP8b6BMg8bROphj54JoqGKSA/VUcnnWKQSf8.jpg?size=200x0&quality=90&crop=14,343,1008,1370&sign=886c6c89aeb653292cb46f11e7d708ac&ava=1
41	42	https://sun1-23.userapi.com/impf/c837621/v837621148/206ed/D58kO_Y1NeU.jpg?size=200x0&quality=90&crop=20,20,564,564&sign=ada99e7ba37a73165f2ae161efa9f3d8&ava=1
42	43	https://sun9-31.userapi.com/c22/u555151/a_a9c9670.jpg?ava=1
43	44	https://sun1-91.userapi.com/impf/c627925/v627925154/2a8d6/6PIdZ3xsYNg.jpg?size=200x0&quality=90&crop=0,0,960,960&sign=9efe55654a78ab5027f81a9b7517b427&ava=1
44	45	https://sun1-29.userapi.com/impf/c625116/v625116158/1670/HhcIZ77Dmfg.jpg?size=200x0&quality=90&crop=26,26,482,754&sign=8996fc523c991f8168e7e8dca068a2ac&ava=1
45	46	https://sun1-88.userapi.com/impf/b_AoGiYuGXSbDH5PKafiMe-waA0YDOI8K_6iVQ/h0TeoCo6X6M.jpg?size=200x0&quality=90&crop=0,32,682,992&sign=7368c253e8b6d2c6cf803b613084348d&ava=1
46	47	https://sun1-87.userapi.com/jecq2LGdyetrymmXcIfcSDZ7145tJPmujcWupQ/TidloAIreMM.jpg?ava=1
47	48	https://sun9-72.userapi.com/c540/u555176/a_6a115e28.jpg?ava=1
48	49	https://sun9-24.userapi.com/c9844/u555179/a_2b82b0bd.jpg?ava=1
49	50	https://sun9-33.userapi.com/c1042/u555185/a_6199263b.jpg?ava=1
50	51	https://sun1-17.userapi.com/impf/c626929/v626929190/3847d/EpMvjm-d9Wc.jpg?size=200x0&quality=90&crop=42,42,876,1196&sign=78653851b260f08ee913119a1f9a4eb8&ava=1
51	52	https://sun1-90.userapi.com/impf/c630217/v630217191/490e0/5dmlOsQNiek.jpg?size=200x0&quality=90&crop=118,228,1313,1885&sign=d21bc36c720289d87f5453f6e3a29c6a&ava=1
52	53	https://sun9-46.userapi.com/c4398/u555201/a_0ddf8a52.jpg?ava=1
53	54	https://sun1-29.userapi.com/impf/c845522/v845522087/1907a8/mvSoUOayKsU.jpg?size=200x0&quality=90&crop=0,0,1620,2160&sign=723c1f07817409e6a6e227b9b82d28a8&ava=1
54	55	https://sun9-36.userapi.com/c9767/u555207/a_29b1c780.jpg?ava=1
55	56	https://sun1-91.userapi.com/9Y31rIDhegLwe7HQQa4a1nyiE17PCjOquGGPCQ/fPvcLDsyYyk.jpg?ava=1
56	57	https://sun1-47.userapi.com/impf/c846322/v846322980/8ee62/MpIrrrnOaLA.jpg?size=200x0&quality=90&crop=3,109,1211,1211&sign=4b00b4d2274069a82882fc3627f06871&ava=1
57	58	https://sun9-28.userapi.com/c1872/u555233/a_35183c90.jpg?ava=1
58	59	https://sun9-23.userapi.com/c9899/u555234/a_631ae9bb.jpg?ava=1
59	60	https://sun1-93.userapi.com/impg/S_B7vj8Eol7AsMVlhTCnfTpEZPIEafE_TAOHhA/YnxOMGNkV5A.jpg?size=200x0&quality=90&crop=1,257,1109,1109&sign=785621415e6ba4199f89a0f08d7d6ebe&ava=1
60	61	https://sun9-33.userapi.com/c09/u555239/a_3f5365b.jpg?ava=1
61	62	https://sun1-25.userapi.com/impg/Y8EuKBeVAu9XRWLsvYgFpD_n-9l6Gv65DY_iaA/3qjqwfgyLUM.jpg?size=200x0&quality=90&crop=71,107,1298,1946&sign=30ce464db7bd92ce0092132ec049586e&ava=1
62	63	https://sun9-18.userapi.com/c15/u555246/a_fc57cee.jpg?ava=1
63	64	https://sun9-8.userapi.com/c10/u555247/a_1ea7dd6.jpg?ava=1
64	65	https://sun1-14.userapi.com/impf/WpjS_RuXQUW-TgLzJyvK0ByZUWcnpjwOb6LrqA/6AdQxZi37Cs.jpg?size=200x0&quality=90&crop=542,118,1547,1547&sign=f9925fc79c366c031f11513112348a66&ava=1
65	66	https://sun1-85.userapi.com/impf/c307307/v307307255/e6c/XfUhepZgI5Y.jpg?size=200x0&quality=90&crop=0,0,750,1000&sign=f8a59c9358677fbf86e7203d0ff0696a&ava=1
66	67	https://sun1-94.userapi.com/impf/c850132/v850132715/a74f9/S8Ik7WWanYg.jpg?size=200x0&quality=90&crop=0,206,546,818&sign=54bf90b150920c7b16fa5089f1a87a75&ava=1
67	68	https://sun1-83.userapi.com/impf/c830608/v830608069/136b6/llNkivIqlLk.jpg?size=200x0&quality=90&crop=442,436,1017,1017&sign=80df33fb23b4b4195dd1255d89a53583&ava=1
68	69	https://sun1-91.userapi.com/impg/c857536/v857536958/115cca/D5Lu2WfSe60.jpg?size=200x0&quality=90&crop=0,89,810,810&sign=5cd503bd3696e0b5de86195c5d1f6fbf&ava=1
69	70	https://sun1-26.userapi.com/impg/J5lIWC9-1xAIKMhx-OumE6qaF-V5RstOQOOixw/Qla4emYVZKw.jpg?size=200x0&quality=90&crop=0,0,1620,1620&sign=13da928badb913d9d566435e90caddd9&ava=1
70	71	https://sun1-22.userapi.com/impf/RKjnQ1Rf5454EneHKeqiNvMkXvmPhu53J1nCvA/IKZtCvhpU-w.jpg?size=200x0&quality=90&crop=108,84,769,769&sign=0ff88ea4f62e7626d645ee38bce92395&ava=1
71	72	https://sun9-62.userapi.com/c11134/u555280/a_2ed0376d.jpg?ava=1
72	73	https://sun1-99.userapi.com/impf/c851420/v851420853/19e0c3/IZgdtdpetb4.jpg?size=200x0&quality=90&crop=0,0,1452,2160&sign=aca8499370bb3a141e6408b4bd5c7a8b&ava=1
73	74	https://sun1-25.userapi.com/impf/c853524/v853524253/38997/HECbHNy2LLI.jpg?size=200x0&quality=90&crop=665,0,1217,1707&sign=327a56422d1db4afbbc09bf5c5834fde&ava=1
74	75	https://sun1-15.userapi.com/impf/cgLFEt3_2-45Wvpsd1Y63tIifcOnxGt3af5zww/xjNoTDZuHDQ.jpg?size=200x0&quality=90&crop=0,0,556,556&sign=0e08cbb76093b9270eacaffbcd9c141f&ava=1
75	76	https://sun1-26.userapi.com/impf/xiUqzm7Eu1q6ZE_Oqt-j3dvY3c-hUxHilNNm7g/w25C5s4oeAg.jpg?size=200x0&quality=90&crop=0,0,600,600&sign=0b33f88fb49c4b1ab8d143ee652d7960&ava=1
76	77	https://sun1-87.userapi.com/impf/c622322/v622322287/1417a/l2hDoDyZd4s.jpg?size=200x0&quality=90&crop=14,5,419,661&sign=379694c556416ed4b7755cea102879f3&ava=1
77	78	https://sun9-68.userapi.com/c9674/u555288/a_722aa89e.jpg?ava=1
78	79	https://sun1-26.userapi.com/impf/c840735/v840735827/2e556/E9xC3w8YBg0.jpg?size=200x0&quality=90&crop=156,33,1751,1751&sign=81978bd357dd246ae62b101ff5bb38b3&ava=1
79	80	https://sun1-15.userapi.com/impf/c848532/v848532968/4f6a4/kDAQeOrRilI.jpg?size=200x0&quality=90&crop=0,0,2048,2048&sign=ba123464d1ede837fc466ea22d17d900&ava=1
80	81	https://sun9-49.userapi.com/c9410/u555309/a_3776f5a6.jpg?ava=1
81	82	https://sun1-88.userapi.com/impf/c841030/v841030312/5df6/k8P11WhqpkA.jpg?size=200x0&quality=90&crop=21,21,598,598&sign=baca701915a42a794add2a27cc875180&ava=1
82	83	https://sun1-25.userapi.com/impf/c840238/v840238460/8d72b/G2ZR2Qn8Ajg.jpg?size=200x0&quality=90&crop=95,24,558,558&sign=0bef2d15360fb168269beee9d64cd8d2&ava=1
83	84	https://sun1-88.userapi.com/impf/c637330/v637330319/b9cf/sITJrR0HUIo.jpg?size=200x0&quality=90&crop=0,0,2160,2160&sign=5151e7dace241a27a019d059f19506cf&ava=1
84	85	https://sun1-18.userapi.com/impf/c625223/v625223320/14c42/Gv5VHXSVZas.jpg?size=200x0&quality=90&crop=402,84,1751,1751&sign=a39a4b7e71f93fe0b5ec2beaa44c95ad&ava=1
85	86	https://sun9-18.userapi.com/c4232/u555321/a_77e20bbe.jpg?ava=1
86	87	https://sun9-25.userapi.com/c10232/u555326/a_e51c41b0.jpg?ava=1
87	88	https://sun1-89.userapi.com/impg/rtNCEm8fniWH1-wzUUGt07fQNBkOb6BYsorE7w/761JwY_V21A.jpg?size=200x0&quality=90&crop=89,143,1477,2017&sign=2f3bfd3c3e37e8d0d9050aeae2fc5279&ava=1
88	89	https://sun1-25.userapi.com/impf/lv5PO9BLg13jTO_nomDMo27pzFj5lbXqHb7H1Q/ApF1CbvW3ps.jpg?size=200x0&quality=90&crop=80,0,480,480&sign=988a28c7fb408bda845bc8da1c4ea64b&ava=1
89	90	https://sun1-96.userapi.com/impg/irX1B4e3dSwNLRLSu_tZFpTRMXrvrLTojQ5E1A/RfcQc8qtMH0.jpg?size=200x0&quality=90&crop=0,361,1437,1437&sign=4ebfb942ba8be50f3780bbd1497749d5&ava=1
90	91	https://sun9-50.userapi.com/c4117/u555340/a_f24e1fcc.jpg?ava=1
91	92	https://sun1-95.userapi.com/impg/iBGv2f4kKD_J7xbnG04c3QZkg94sg9q-l5UWAA/KbIMYIu8voo.jpg?size=200x0&quality=90&crop=0,106,1344,1344&sign=47bc4ad90d734efb44ce229f04cbd085&ava=1
92	93	https://sun1-29.userapi.com/impf/c848732/v848732221/1589cd/F-_SX9vGfWk.jpg?size=200x0&quality=90&crop=202,0,809,809&sign=b38ec46dc07acf65d533983c75c626c9&ava=1
93	94	https://sun1-47.userapi.com/impf/aGUu4xfmTW_UeujZOJKS0X8EqAsh9XKiuVgF1g/vUzYn0AOnDg.jpg?size=200x0&quality=90&crop=432,0,1759,1920&sign=f558b1c17485ff3784686fb2c5a339fd&ava=1
94	95	https://sun9-17.userapi.com/c592/u555354/a_f5547aa9.jpg?ava=1
95	96	https://sun1-94.userapi.com/impg/ClmvPrGRhXe_Yfq7TTV1KCClTvSH5stN67GJMA/sXC3yNdGeF8.jpg?size=200x0&quality=90&crop=35,35,739,1009&sign=0fe0c7dc4e718c98161165a29549c03a&ava=1
96	97	https://sun1-84.userapi.com/impf/c836722/v836722366/454ed/4RWuGcRJ9ZA.jpg?size=200x0&quality=90&crop=0,0,2045,2048&sign=cad3084b349ad1f4d40f593b0744ea25&ava=1
97	98	https://sun1-23.userapi.com/impf/c824201/v824201111/abe5d/UnlQJxT9vPU.jpg?size=200x0&quality=90&crop=0,45,962,962&sign=e295759b75642bc51f2293c35f292566&ava=1
98	99	https://sun1-19.userapi.com/impf/c849324/v849324057/16484c/UOyq6yPmcVs.jpg?size=200x0&quality=90&crop=250,91,370,370&sign=084e0bedac37db20df4f9d083a1400a4&ava=1
99	100	https://sun1-90.userapi.com/impf/c636526/v636526376/58442/FJfhwtjRRDs.jpg?size=200x0&quality=90&crop=20,20,560,560&sign=62b3462e2e6e0cc897e5593530f73a9d&ava=1
100	101	https://sun1-96.userapi.com/impg/RzrDk9BjFF_R9xQHdoVLpNlAv2MKBtw_15CNzA/ELM1bUFJIfY.jpg?size=200x0&quality=90&crop=195,1,806,806&sign=43b8ea757aa33d6390d573757efc2feb&ava=1
101	102	https://sun1-28.userapi.com/impf/c631624/v631624380/1ebaf/PAJ4glwWA5M.jpg?size=200x0&quality=90&crop=0,0,629,629&sign=235760edcbf3856b1902e81707936d9a&ava=1
102	103	https://sun1-30.userapi.com/impg/-M-xYAmMk3dhuqVGcM7secProjLN6ucmesQc1w/85kjwCd2UmA.jpg?size=200x0&quality=90&crop=0,28,600,601&sign=1cb3ab0bb9d190bd7a387f679a1fc31f&ava=1
103	104	https://sun1-99.userapi.com/impf/c855124/v855124978/1668e3/W_LqQw--Pz8.jpg?size=200x0&quality=90&crop=0,0,1079,1079&sign=704792a5901dc6c8e4786fe2513b185d&ava=1
104	105	https://sun9-28.userapi.com/c16/u555388/a_c323cdf.jpg?ava=1
105	106	https://sun1-21.userapi.com/impf/c604425/v604425390/38939/mYNKfS6Zj3k.jpg?size=200x0&quality=90&crop=63,63,1793,1793&sign=455edb675d91c3c8d75088efc6aded1a&ava=1
106	107	https://sun1-23.userapi.com/impf/EW7EuOteuI6pbrYDwWzXUsp_AAt_ZA1Zc0oxVg/vst9VS42zRQ.jpg?size=200x0&quality=90&crop=0,0,1200,1600&sign=bcfcef8723bbac38ac133b748336cc86&ava=1
107	108	https://sun9-7.userapi.com/c16/u555399/a_110f099.jpg?ava=1
108	109	https://sun1-97.userapi.com/impg/yW4bhzqYjpf3zFMNCxyXdp4v7CAkYroAqPPoKg/nQXCpDiwJvM.jpg?size=200x0&quality=90&crop=135,247,295,295&sign=db03d00bb91a21bd96668df8d0d9a3f1&ava=1
109	110	https://sun1-92.userapi.com/impg/c858520/v858520023/1de37d/mjLQT5X_fTg.jpg?size=200x0&quality=90&crop=0,350,1620,1620&sign=f4f83d7d5edfb20ea1a80969542e6861&ava=1
110	111	https://sun9-19.userapi.com/c32/u555409/a_63a7acb.jpg?ava=1
111	112	https://sun9-20.userapi.com/c08/u555414/a_e7e7815.jpg?ava=1
112	113	https://sun9-40.userapi.com/c9817/u555415/a_44f13fea.jpg?ava=1
113	114	https://sun9-59.userapi.com/c16/u555421/a_e16b2f3.jpg?ava=1
114	115	https://sun1-18.userapi.com/impf/c625322/v625322423/40fb3/WJyH9Nbc8Ds.jpg?size=200x0&quality=90&crop=0,80,720,1079&sign=cd30a1232b1b57537ea061418019adb2&ava=1
115	116	https://sun1-89.userapi.com/impf/c638128/v638128436/2cff6/xuGgEpIOS_Q.jpg?size=200x0&quality=90&crop=180,0,719,719&sign=001437b006c57cc4f48e4feb4d30b840&ava=1
116	117	https://sun9-29.userapi.com/c308617/u555441/a_d1d73a74.jpg?ava=1
117	118	https://sun1-19.userapi.com/impf/c638727/v638727445/26219/GdYo7O5RINc.jpg?size=200x0&quality=90&crop=224,7,370,370&sign=96a6d0fffcb79516c417e27403e1a23a&ava=1
118	119	https://sun1-23.userapi.com/impf/c624330/v624330446/37829/Kv04elaqCSc.jpg?size=200x0&quality=90&crop=0,0,960,960&sign=a951364c8a37d206dfc8fc89f33544a7&ava=1
119	120	https://sun9-32.userapi.com/c10691/u555448/a_ed82464c.jpg?ava=1
120	121	https://sun1-25.userapi.com/impg/p32hhj-JIleBqueh6AxQAi2qnBBdKUlgULgIkA/fXADSSFbxzQ.jpg?size=200x0&quality=90&crop=312,0,1919,1919&sign=2bb26d7b903322e1cdbc44ef198513a4&ava=1
121	122	https://sun1-20.userapi.com/impg/8EbSk45GEamYz35FcZVoiPLAasfBDUIUdCg5PA/_vxztrSt0Rc.jpg?size=200x0&quality=90&crop=4,182,950,950&sign=c25879671194f61867d011b1582bb047&ava=1
122	123	https://sun1-14.userapi.com/impf/c637124/v637124870/8b649/6X4PMdOZrQ4.jpg?size=200x0&quality=90&crop=0,248,639,639&sign=5bd4bb76f9d50ea9a0d9df374208cc7d&ava=1
123	124	https://sun1-29.userapi.com/impf/c856120/v856120337/d7f68/AyZLn7DPCR0.jpg?size=200x0&quality=90&crop=0,302,1537,1537&sign=6a8351437900597805a312d0df17a040&ava=1
124	125	https://sun9-51.userapi.com/c10/u555467/a_91a041d.jpg?ava=1
125	126	https://sun1-25.userapi.com/impf/4XMghvh6fNUHm9uAMRItwWqoXmDSZmkJv35NlA/XeE3UkM-Lwo.jpg?size=200x0&quality=90&crop=0,0,720,720&sign=0fcd515b05ec545a8fb5acee683cca2a&ava=1
126	127	https://sun1-99.userapi.com/impf/c852120/v852120422/180bca/iDG8o6l1j1Q.jpg?size=200x0&quality=90&crop=135,0,809,809&sign=bbbe0d2cfe13fc2428c764e2367c0256&ava=1
127	128	https://sun1-87.userapi.com/impg/nhNB6PTdawPxxe-TubLOrajk2p0j3_yaoTFvCA/aNwP2xg8wEQ.jpg?size=200x0&quality=90&crop=0,338,1620,1620&sign=2471af0e5b471240127a0107b653433f&ava=1
128	129	https://sun1-99.userapi.com/impf/c855224/v855224569/e06f2/NfUfI3SB0QM.jpg?size=200x0&quality=90&crop=254,0,1528,1528&sign=1f51cee1dc9c3067eccf6787ff1d4baf&ava=1
129	130	https://sun1-96.userapi.com/impf/lDgo1EoOoTRJ0HQJsDF9FLcgR4G1O3E963OkSw/BAVj3YkwXvg.jpg?size=200x0&quality=90&crop=250,0,679,679&sign=fd7f7c081d6981a5d93fba6dac0b5c1d&ava=1
130	131	https://sun1-84.userapi.com/impf/67CSQaSohLM09ycQqLK3sr1c9Ij70nc808dg2A/FIwilRVBIs8.jpg?size=200x0&quality=90&crop=0,0,451,604&sign=8a776f08dbe6f43f621cd653f7fcbaf0&ava=1
131	132	https://sun9-32.userapi.com/c9674/u555514/a_ef5d4fd3.jpg?ava=1
132	133	https://sun9-14.userapi.com/c439/u555517/a_e4b5ceef.jpg?ava=1
133	134	https://sun1-95.userapi.com/impf/c637120/v637120518/1223f/eyBBwM-KB_Q.jpg?size=200x0&quality=90&crop=62,1,512,512&sign=3c61c0403d6c13b7f3c07809d5d7491e&ava=1
134	135	https://sun9-8.userapi.com/c9288/u555520/a_c2382e6a.jpg?ava=1
135	136	https://sun1-47.userapi.com/impf/c625521/v625521526/7e4/VfmkCyGbRO4.jpg?size=200x0&quality=90&crop=0,0,960,960&sign=5c820adcb42657c301113394fd705106&ava=1
136	137	https://sun9-65.userapi.com/c313/u555530/a_f497bcd2.jpg?ava=1
137	138	https://sun1-14.userapi.com/impf/c623619/v623619536/4c80e/7cNBAbqw3s8.jpg?size=200x0&quality=90&crop=0,33,1365,1991&sign=78d6f3b438f397fb260a841a5974f982&ava=1
138	139	https://sun1-88.userapi.com/impf/c633130/v633130540/55c9/ZkbMZNDc35w.jpg?size=200x0&quality=90&crop=6,0,488,488&sign=bbdaab6f245a95d5740abe3c3f3c9eef&ava=1
139	140	https://sun1-95.userapi.com/impf/c625622/v625622552/2545f/UbvoDrERFiw.jpg?size=200x0&quality=90&crop=20,20,437,524&sign=eda1d6fd1a914f1a328f32c046628648&ava=1
412	413	https://sun9-25.userapi.com/c304403/u556540/a_cef085ee.jpg?ava=1
140	141	https://sun1-14.userapi.com/impf/c624523/v624523555/4bdcd/7w3GmoHuYrw.jpg?size=200x0&quality=90&crop=241,0,609,649&sign=2968c70f3a336f3d88d1ff6334018172&ava=1
141	142	https://sun1-83.userapi.com/impg/svxSP9TieGOmgv_M-abrlWcXfCLC06ZPhkow9Q/KEcbfQ3MXtQ.jpg?size=200x0&quality=90&crop=1,1,1197,1197&sign=3023b0921f7c7793a1997ec204a65cae&ava=1
142	143	https://sun1-99.userapi.com/impf/c630718/v630718557/3cab9/gK8020b5XeA.jpg?size=200x0&quality=90&crop=0,239,607,607&sign=5631223dbc5b7e99ae24e58e2703624f&ava=1
143	144	https://sun9-74.userapi.com/c9952/u555561/a_954eaa34.jpg?ava=1
144	145	https://sun9-60.userapi.com/c16/u555564/a_50a80df.jpg?ava=1
145	146	https://sun9-19.userapi.com/c204/u555565/a_b9a9e524.jpg?ava=1
146	147	https://sun1-94.userapi.com/impg/c857536/v857536582/13f5e9/rYk8mXDkrhY.jpg?size=200x0&quality=90&crop=0,779,1152,1152&sign=33d73b3b898d2d1d34d752f564298ab6&ava=1
147	148	https://sun1-93.userapi.com/impf/c308119/v308119571/9864/JQRtqpTD3Yg.jpg?size=200x0&quality=90&crop=254,42,811,811&sign=40b0dd4e9e7b098c26e312557021ecf0&ava=1
148	149	https://sun1-28.userapi.com/impf/c857732/v857732703/b33dd/vXXwyRGMllg.jpg?size=200x0&quality=90&crop=4,4,803,803&sign=747272b530d95b3e5be463f582d58374&ava=1
149	150	https://sun1-84.userapi.com/impf/nPIXTZo0h_QO_M5QKXdsAQjbkkBODxVKVj0F7Q/tPqMyjrnHtc.jpg?size=200x0&quality=90&crop=0,0,1136,1136&sign=eadcfd77b6f702978df26d009674c5e8&ava=1
150	151	https://sun1-23.userapi.com/impg/c858036/v858036119/104308/p6i-q6ATlaQ.jpg?size=200x0&quality=90&crop=481,990,370,370&sign=deb75a11c1efa30a1d5bd0f0164f09b6&ava=1
151	152	https://sun1-92.userapi.com/impf/FPKavWFc6TzQqGW2aGHTa1SXJ955dksb5oWcaQ/wABHnX5KZkM.jpg?size=200x0&quality=90&crop=678,0,1700,1701&sign=6bb94ff7d7d8cf7d8b6d3ba38b21d694&ava=1
152	153	https://sun9-16.userapi.com/c271/u555598/a_1940dfad.jpg?ava=1
153	154	https://sun9-14.userapi.com/c16/u555603/a_3048b1c.jpg?ava=1
154	155	https://sun1-88.userapi.com/impf/c855320/v855320926/d9aa1/zfEhrTlqzYI.jpg?size=200x0&quality=90&crop=160,0,959,959&sign=bf27e1e2741ee7326856ad34c73a9df2&ava=1
155	156	https://sun1-96.userapi.com/impg/fKbV3CTTnaC1zjwC1pGcISTHKzWJsJy9zx9yqQ/85n1-xYUwE4.jpg?size=200x0&quality=90&crop=2,277,1607,1607&sign=5e9b94d200b1f18a51137cfeba2a9b9d&ava=1
156	157	https://sun1-83.userapi.com/impf/Zagy849n5UrHYLM2wMyYGxKoQ6Ma-nrH1w94Fg/Q4fKEfRMhRw.jpg?size=200x0&quality=90&crop=495,0,1440,1440&sign=28f995d410818f974c2ab7162c6f0647&ava=1
157	158	https://sun1-17.userapi.com/impf/c636819/v636819627/1a2c5/lmOJajTq2us.jpg?size=200x0&quality=90&crop=0,0,1132,1132&sign=dbdbac64ab2f13d04e7435d59392c535&ava=1
158	159	https://sun1-25.userapi.com/impf/c847120/v847120292/bb9af/8SJC9A27TKo.jpg?size=200x0&quality=90&crop=0,0,500,678&sign=ffd4659b50d974add88b95dc1115eb9f&ava=1
159	160	https://sun9-69.userapi.com/c10/u555632/a_0573b29.jpg?ava=1
160	161	https://sun1-89.userapi.com/impf/c857232/v857232557/402e/3vEBgFutPeU.jpg?size=200x0&quality=90&crop=157,160,1410,1653&sign=c354111f7ef0daabf10fea73e48c3b32&ava=1
161	162	https://sun1-20.userapi.com/impf/I0yIoNmBP4PAjKJfuZn55Rz_fe-S5yz3u14FRQ/o2lD3gn3jJI.jpg?size=200x0&quality=90&crop=563,0,1441,1441&sign=9640c2e28f23d1a2e4d9301c147ce751&ava=1
162	163	https://sun9-15.userapi.com/c5369/u555638/a_570ca422.jpg?ava=1
163	164	https://sun1-14.userapi.com/impf/c858324/v858324476/1a328/NiFkp7zSzBU.jpg?size=200x0&quality=90&crop=128,425,905,1352&sign=8dd49b2d3061d4c896ee70b119ce0a55&ava=1
164	165	https://sun9-45.userapi.com/c5593/u555651/a_e50aff38.jpg?ava=1
165	166	https://sun1-87.userapi.com/impg/eSPwFTzk_gyn1ej_QXfNnN_vrz0d8YWiqnAEdw/CO64N8S8WY4.jpg?size=200x0&quality=90&crop=2,2,2155,2155&sign=49b761847626ddfc4de546d90633e960&ava=1
166	167	https://sun1-97.userapi.com/impf/c622430/v622430663/41078/nwVIsfoW_RM.jpg?size=200x0&quality=90&crop=0,0,851,852&sign=28a0c87d1d0ebb420b5e9887d6a257cc&ava=1
167	168	https://sun1-47.userapi.com/impf/c840435/v840435955/63a64/7Sqv0sRSSS8.jpg?size=200x0&quality=90&crop=549,66,1301,1045&rotate=270&sign=b034f995d16c8ad61db10aa9f1ed0172&ava=1
168	169	https://sun1-90.userapi.com/impf/c623924/v623924677/f3c4/N_2EHWXv3KI.jpg?size=200x0&quality=90&crop=546,169,1538,1538&sign=11328aa885998b79ce03e6685648f347&ava=1
169	170	https://sun9-25.userapi.com/c16/u555685/a_1e10b1f.jpg?ava=1
170	171	https://sun9-64.userapi.com/c11/u555688/a_f18cb97.jpg?ava=1
171	172	https://sun9-70.userapi.com/c29/u555690/a_9ba7f85.jpg?ava=1
172	173	https://sun1-84.userapi.com/impf/c625219/v625219696/42bcc/C3ciZrJ343E.jpg?size=200x0&quality=90&crop=24,24,701,701&sign=dc85507dd5157fd9023d86875263787a&ava=1
173	174	https://sun1-17.userapi.com/impf/c824501/v824501112/76ee6/Pro28_eZYG4.jpg?size=200x0&quality=90&crop=0,0,1537,1537&sign=30638e794c86fca71bfb1cefa534ca35&ava=1
174	175	https://sun1-98.userapi.com/impf/c631316/v631316702/51f8/cjeuxV0Wv1U.jpg?size=200x0&quality=90&crop=187,0,2048,2048&sign=f4b57e23c430e801cda4b131cd3a71b6&ava=1
175	176	https://sun1-84.userapi.com/impf/c5143/v5143704/1070/KqJKEVUi9hw.jpg?size=200x0&quality=90&crop=199,40,878,878&sign=186418ac6a12e99c5b4937e5ef8cd0a2&ava=1
176	177	https://sun1-84.userapi.com/impf/6BOWpqE0a1cN1wW0lfiMf2PZwvyLzYCoJs8S3Q/S1DLCbLS-5Y.jpg?size=200x0&quality=90&crop=105,0,407,407&sign=58725293646d566d6d6d9bbab40cb0df&ava=1
177	178	https://sun1-47.userapi.com/impf/c841223/v841223484/11335/QCwQXVpwUv8.jpg?size=200x0&quality=90&crop=143,0,809,809&sign=0916b0c2554f8a151e800e74267b75cb&ava=1
178	179	https://sun1-14.userapi.com/impf/c844417/v844417514/14afb1/CIv6x3iXw1w.jpg?size=200x0&quality=90&crop=0,60,739,1009&sign=e39163dca5d0bc71465bfbe5115c3e8e&ava=1
179	180	https://sun1-17.userapi.com/impf/c854524/v854524390/e41cc/FvLgV1aiiIs.jpg?size=200x0&quality=90&crop=10,16,2132,2132&sign=c7463b9a46e24435439e80b86a260079&ava=1
180	181	https://sun9-28.userapi.com/c698/u555717/a_d810b481.jpg?ava=1
181	182	https://sun1-99.userapi.com/impg/xbF7_N04pbd8-CQdOyyOoHalsbjOS6x66awLhg/iz-CdkpQ2UQ.jpg?size=200x0&quality=90&crop=6,328,1420,1421&sign=bf516ce1e30884ea4c5dedc55cf0bbcc&ava=1
182	183	https://sun1-17.userapi.com/impf/c631430/v631430723/94d0/YRz7LIAoOW8.jpg?size=200x0&quality=90&crop=259,0,960,960&sign=dbed8fc3528a73398eda5e049b112901&ava=1
183	184	https://sun1-20.userapi.com/impf/c849036/v849036361/cf298/A0PS0xyqggk.jpg?size=200x0&quality=90&crop=0,54,758,758&sign=ae32d833cbf6fa44fbfffc682651487f&ava=1
184	185	https://sun1-23.userapi.com/impg/c858424/v858424635/21ddd1/3DhINLGdo4A.jpg?size=200x0&quality=90&crop=82,0,1073,1610&sign=0bba0d10715fd14dfcb1cb4eea33a487&ava=1
185	186	https://sun1-84.userapi.com/impf/c846523/v846523549/1d7eca/c3tHmyVJ-o0.jpg?size=200x0&quality=90&crop=928,0,958,1433&sign=2230fb93ab5edfd54e842526e50dea69&ava=1
186	187	https://sun1-85.userapi.com/impf/c633527/v633527741/fd2a/MzcPgHAKydg.jpg?size=200x0&quality=90&crop=320,0,1920,1920&sign=37b4b5592255876f74b22a485622ca37&ava=1
187	188	https://sun1-29.userapi.com/impf/c850428/v850428859/204ac/kajS6b6byS0.jpg?size=200x0&quality=90&crop=0,242,720,720&sign=75ccc43dfd6f5a8e5eb9fdf0d1c6a045&ava=1
188	189	https://sun9-23.userapi.com/c4161/u555744/a_27a56901.jpg?ava=1
189	190	https://sun1-87.userapi.com/impf/c849124/v849124231/cabbf/XwzAQcp-pJM.jpg?size=200x0&quality=90&crop=542,59,1538,1538&sign=90fdbf8377e27c3a3d618ac7c9ca922d&ava=1
190	191	https://sun1-99.userapi.com/impf/6tsej1gBOoll51O2oG4JgnBAA8791uq_PC6GtA/Rv3a3HergHU.jpg?size=200x0&quality=90&crop=67,67,1401,1913&sign=b2a80d5cc59a531bb2f94608981a02cd&ava=1
191	192	https://sun1-86.userapi.com/impf/c626328/v626328773/2f7df/GzEiAXudKGk.jpg?size=200x0&quality=90&crop=286,1,1057,1057&sign=3b6517dd0ce3e36bc399786723e7289b&ava=1
192	193	https://sun1-22.userapi.com/impf/Yi8S98gd0H_8dsHcNy60ph3dKMgGTdb-tDmnkg/xctAWNSY7b0.jpg?size=200x0&quality=90&crop=508,0,1713,1714&sign=cfd741f57ee0aad7b21cf1841cc08ced&ava=1
193	194	https://sun1-96.userapi.com/impf/c850036/v850036694/ae930/mFcLcww3_EQ.jpg?size=200x0&quality=90&crop=0,0,1728,2160&sign=f78e4fc241f6ad43f6056803ed039b1d&ava=1
194	195	https://sun1-30.userapi.com/impg/wvEuUhx2HLRYsBRjPUWjGtou8CLgvoSZxOHzxw/Nyw3c6QxiHs.jpg?size=200x0&quality=90&crop=0,135,809,809&sign=fa5e5345eb8aaf7e67e9d1fc05032f85&ava=1
195	196	https://sun1-91.userapi.com/impf/c639323/v639323259/3a45c/2tDSF0h3ZKI.jpg?size=200x0&quality=90&crop=121,207,1298,1946&sign=4a4f1f652ff10a1e3d7623c9447ef37e&ava=1
196	197	https://sun1-18.userapi.com/impf/c856032/v856032218/e6166/klFbyeZAgsE.jpg?size=200x0&quality=90&crop=406,0,458,684&sign=23ee8729507ebfe71ed607388f95f3e2&ava=1
197	198	https://sun1-28.userapi.com/impf/LdnfgycRnlHzQwZl5FcmwBVMSe30PTfdLGiBLA/v-RVcNqmpu4.jpg?size=200x0&quality=90&crop=0,0,1365,2048&sign=c35b06b19b6db16338b1b35a77ed51d1&ava=1
198	199	https://sun1-89.userapi.com/impg/c858232/v858232486/1eec3c/tfIDt-N0s0E.jpg?size=200x0&quality=90&crop=0,87,434,434&sign=2b6824e31ab98d093491d34665de13d5&ava=1
199	200	https://sun1-95.userapi.com/impf/c630129/v630129808/414d6/N5x1_Le8OQY.jpg?size=200x0&quality=90&crop=0,0,480,640&sign=51c95baa27ef4f5656b48d6ddac1a27e&ava=1
200	201	https://sun9-18.userapi.com/c4567/u555809/a_e02b1d6a.jpg?ava=1
201	202	https://sun9-8.userapi.com/c4175/u555816/a_256d6ac4.jpg?ava=1
202	203	https://sun9-54.userapi.com/c9587/u555818/a_50cbcd14.jpg?ava=1
203	204	https://sun1-17.userapi.com/impf/c629300/v629300819/61aef/jnHtfvJWAsU.jpg?size=200x0&quality=90&crop=0,161,639,639&sign=c0ecd27404f9cddc0b84b82397922f3e&ava=1
204	205	https://sun1-14.userapi.com/impf/c622716/v622716823/4a2cd/DEQ9nb7BVa0.jpg?size=200x0&quality=90&crop=125,0,354,456&sign=883e908b422c920f08f517824f9ec411&ava=1
205	206	https://sun1-97.userapi.com/impf/c622823/v622823677/2b366/SIgk-31W9TY.jpg?size=200x0&quality=90&crop=430,82,649,661&sign=a229deee56332f49093eac648ecf956d&ava=1
206	207	https://sun1-96.userapi.com/impf/c626817/v626817827/478ef/yN_05q2zSQE.jpg?size=200x0&quality=90&crop=0,0,200,299&sign=b2eb8a0bbb8d07daf9d2ff93d179f0f0&ava=1
207	208	https://sun9-8.userapi.com/c4143/u555829/a_6c3dd908.jpg?ava=1
208	209	https://sun1-30.userapi.com/impf/xDdiW30Lb1tk66eRbCRuIx4qCufBvMVllDXHyA/vKXwN62yMV8.jpg?size=200x0&quality=90&crop=320,0,1920,1920&sign=479aae37d3109b92398a3840414605c6&ava=1
209	210	https://sun1-84.userapi.com/impf/c638825/v638825737/55fc8/3PYwX5rAebA.jpg?size=200x0&quality=90&crop=0,0,640,640&sign=f712396dc18b1734169648380ff642c8&ava=1
210	211	https://sun1-27.userapi.com/impg/Vgt3Uv1VVWtNfqopL4o1nU3ayzlAyCu0AZcIRg/_HuJcJ2d8YY.jpg?size=200x0&quality=90&crop=0,162,960,961&sign=3ae3dad628718a55e9759bbb2f116e65&ava=1
211	212	https://sun1-96.userapi.com/impf/c622517/v622517848/15fb1/vJ5m9y3wAvA.jpg?size=200x0&quality=90&crop=495,84,1285,1286&sign=6197a1ff06b4cf53f8bdd7892f3fa376&ava=1
212	213	https://sun9-6.userapi.com/c9269/u555849/a_7dc9a2c5.jpg?ava=1
213	214	https://sun9-11.userapi.com/c10299/v10299857/3ec/8tEd2I6pRlU.jpg?ava=1
214	215	https://sun1-21.userapi.com/impg/jSbVVK1HjqujwtoWGi5T686uwlAIvbZFsP3XEA/XkRQljlBnJo.jpg?size=200x0&quality=90&crop=1,1,1427,1427&sign=d40ec431a357a365d51f364ffb165176&ava=1
215	216	https://sun1-89.userapi.com/impf/c824701/v824701710/8a960/wr2imHeU6H8.jpg?size=200x0&quality=90&crop=0,0,449,600&sign=a92703c3eea18245f3c603a43ab7e918&ava=1
216	217	https://sun1-85.userapi.com/impf/c840530/v840530639/4bd7f/HB-TjfcKis0.jpg?size=200x0&quality=90&crop=2,308,1616,1616&sign=8f24885e4cea6312be539e1406320a5d&ava=1
217	218	https://sun1-99.userapi.com/impf/c631930/v631930867/2e673/ZqN1ZzTP7r8.jpg?size=200x0&quality=90&crop=0,0,1738,2160&sign=d7b4ce3d7d8a59a84debe4482d13519f&ava=1
218	219	https://sun1-20.userapi.com/impf/c837721/v837721116/57805/XAkIQSCBTFw.jpg?size=200x0&quality=90&crop=161,16,836,836&sign=0da5dea37f4f61ae65bb2a8915227dd4&ava=1
219	220	https://sun1-84.userapi.com/impf/c851424/v851424492/1cd178/msncmoGesxg.jpg?size=200x0&quality=90&crop=0,0,1620,2160&sign=bb0f663f37adebb4b819e97446f2633a&ava=1
220	221	https://sun1-98.userapi.com/impg/zh2Lvzx350Ys1bK5-G2JL1TGtLxjw61PhGZ9YA/UKBllBP7mV8.jpg?size=200x0&quality=90&crop=28,28,584,798&sign=c08a2017768237d316edaf44103b44d2&ava=1
221	222	https://sun1-22.userapi.com/impg/hZ-xXsI0O-oqlPtDw6v9-_p8HyXz88erhxEkDg/qghHDSf_gHA.jpg?size=200x0&quality=90&crop=0,86,864,864&sign=52e7da73bddb5279ff867e133ef60812&ava=1
222	223	https://sun1-92.userapi.com/impf/c852024/v852024379/59aa2/d-HKMAo0jok.jpg?size=200x0&quality=90&crop=0,97,1242,1242&sign=92c155230859daa50dc8dd84463520e3&ava=1
223	224	https://sun1-22.userapi.com/impg/c853520/v853520745/24a84e/puPnjsw4dDU.jpg?size=200x0&quality=90&crop=0,0,640,1136&sign=700f74c04e3a1f68b9a5366f08c9b60e&ava=1
224	225	https://sun1-20.userapi.com/impf/c851224/v851224035/19423e/JxxVxE78Hws.jpg?size=200x0&quality=90&crop=0,92,960,960&sign=89cb8ac8ab58243e2bdb51333cca86b2&ava=1
225	226	https://sun9-64.userapi.com/c08/u555898/a_78f9f63.jpg?ava=1
226	227	https://sun1-18.userapi.com/impf/c637122/v637122899/4b3ca/1rhOMV954pY.jpg?size=200x0&quality=90&crop=593,4,1221,1703&sign=9fafcf979ff2c8341b795f3863345d42&ava=1
227	228	https://sun1-91.userapi.com/impg/FYnh-V0154itYBoLR-E_gCdbMoGXKNgpaG2jqw/PufLVaT6CvY.jpg?size=200x0&quality=90&crop=79,173,860,1287&sign=6dc2f1e252855d178b7b4d17e21647d6&ava=1
228	229	https://sun9-35.userapi.com/c09/u555902/a_db69388.jpg?ava=1
229	230	https://sun1-26.userapi.com/impg/GwUTNUjXFiubCnU1WIUGsNATv0eyTjpmMO1b9w/xW1VV6tn1f4.jpg?size=200x0&quality=90&crop=607,88,1153,1153&sign=5b67e0156a8ae31411c62e334d09b6d1&ava=1
230	231	https://sun1-94.userapi.com/impg/_mYYFcWp2oe-Efq4dPU6rm2aM6boyNZCLD5EpA/gkRfxmk8hLA.jpg?size=200x0&quality=90&crop=0,0,749,838&sign=fc304c2aa040dbde5e75d2264d1040cb&ava=1
231	232	https://sun1-88.userapi.com/impf/xaE1IGmZ0k7tRoxjT7JsFFxmZ4Po-ZjwJAgVOw/dqN7HPF3cfw.jpg?size=200x0&quality=90&crop=30,57,1337,1913&sign=3d30c2fb11f99cc1015583512580ca50&ava=1
232	233	https://sun1-92.userapi.com/impg/wviHhI0e_DMFvNr1f-IpNqzFt3cGHDAX_vZ9OA/G8UyRz0sT4o.jpg?size=200x0&quality=90&crop=20,20,434,555&sign=ecd9ae1dacaa31772a1a847a777ca41a&ava=1
233	234	https://sun9-24.userapi.com/c10/u555918/a_86df513.jpg?ava=1
234	235	https://sun1-96.userapi.com/impf/c844321/v844321515/11eb46/qmqeg1M7PZg.jpg?size=200x0&quality=90&crop=0,0,1364,1500&sign=2df3d90479044e94a15464389813b0dd&ava=1
235	236	https://sun9-42.userapi.com/c10856/u555928/a_5633f6c9.jpg?ava=1
236	237	https://sun9-20.userapi.com/c9698/u555929/a_71d2db93.jpg?ava=1
237	238	https://sun9-29.userapi.com/c5874/u555932/a_717c3aef.jpg?ava=1
238	239	https://sun1-16.userapi.com/impf/c624718/v624718946/4b188/KacI3eG34hE.jpg?size=200x0&quality=90&crop=0,0,959,959&sign=eb3c3b22a04e24afa824bc61bcba23bb&ava=1
239	240	https://sun1-19.userapi.com/impf/c846418/v846418598/141d4/urKHwfVQI_c.jpg?size=200x0&quality=90&crop=0,473,1213,1213&sign=cff56b6f02e460f1fcafbe748c13770b&ava=1
240	241	https://sun1-90.userapi.com/hzks6GYiMtcyv13m2RcIUkSdQyDa-YQNqC9Qow/cCCTpfdM5Bg.jpg?ava=1
241	242	https://sun1-26.userapi.com/impf/YsVATF1CBMQ_oy7GwfwbJd4613ICSDGuCvFSUw/mwlkGz75qV0.jpg?size=200x0&quality=90&crop=402,142,327,410&sign=1f005d3d1b07a3ea04c2d940ff1db4d9&ava=1
242	243	https://sun1-29.userapi.com/impf/c855728/v855728200/138963/JB4FYjsAdc0.jpg?size=200x0&quality=90&crop=160,0,960,960&sign=eef57d8aa18181460adddbd8ef101d79&ava=1
243	244	https://sun1-29.userapi.com/impf/c852028/v852028270/4ea03/dUw01sZlcVs.jpg?size=200x0&quality=90&crop=627,0,1114,1114&sign=2f521027adfeb1b702fcfff6f3d8a06c&ava=1
244	245	https://sun1-87.userapi.com/impf/c849024/v849024841/126627/H45ygBk7Kvw.jpg?size=200x0&quality=90&crop=335,0,1920,1920&sign=6d4932c29ad136c5f5bb05dea90fd202&ava=1
245	246	https://sun9-13.userapi.com/c704/u555961/a_865436b6.jpg?ava=1
246	247	https://sun1-98.userapi.com/impf/c628526/v628526963/7321/r9gZOqR966Q.jpg?size=200x0&quality=90&crop=0,0,959,959&sign=0b21c12dca203864c03cfb6d357eb259&ava=1
247	248	https://sun1-99.userapi.com/impf/c850620/v850620273/12aee4/HYJpDPBQca0.jpg?size=200x0&quality=90&crop=0,0,1440,2160&sign=a0ae1f36076ac75111a5df3d489ed39c&ava=1
248	249	https://sun9-3.userapi.com/c15/u555972/a_f9a6279.jpg?ava=1
249	250	https://sun1-20.userapi.com/impf/3Sop8Llnf_9Iqix8pCjy3NqARqHmG8FwhVABfw/pXmRAAl2UNM.jpg?size=200x0&quality=90&crop=0,246,806,806&sign=25199c7d5c593efbb95200f9a8af62fd&ava=1
250	251	https://sun1-19.userapi.com/impf/oRUp7uPRoPr_DtKU0l2ge70_c-8CR9-ylZX0gw/ajODn5ht7nc.jpg?size=200x0&quality=90&crop=436,82,1608,1608&sign=17bb253c5fa95a68969f7af6b5030758&ava=1
251	252	https://sun1-30.userapi.com/impg/KANW7YGthYJsDwCo1-O07WAfx7WCqJNMJMw7UQ/8A8qlxZF0SE.jpg?size=200x0&quality=90&crop=0,187,1104,1104&sign=ca0de9748628dffc05bd272cabe6d7ee&ava=1
252	253	https://sun9-1.userapi.com/c27/u555982/a_ce5e3e8.jpg?ava=1
253	254	https://sun1-17.userapi.com/impg/unQdJKpc246E0iCgCaHmc25gziVOWWYXSIHO9g/tsPDbiX95mE.jpg?size=200x0&quality=90&crop=3,9,587,587&sign=3dda67cd7069443282cd6a1ba2e2caa0&ava=1
254	255	https://sun1-84.userapi.com/impf/c834403/v834403340/178e7a/HHL9QPh2kng.jpg?size=200x0&quality=90&crop=383,84,769,769&sign=3f3f9e8aa1eca2b8d09b844b6cd8c501&ava=1
255	256	https://sun1-85.userapi.com/impf/c852032/v852032089/1b56d1/Se-z1_2fhNA.jpg?size=200x0&quality=90&crop=69,0,299,299&sign=55cd2b8782c6c83430544a5cf4bf87d8&ava=1
256	257	https://sun1-20.userapi.com/impg/VtRk3TZAox2PODBlNyUgeAzSNfA4Prh3FTHCcA/M5JQyRlrDKo.jpg?size=200x0&quality=90&crop=0,63,450,450&sign=93867428f87dfc5a4c36c4c2f25fb0d3&ava=1
257	258	https://sun9-70.userapi.com/c9728/u555998/a_1ee00118.jpg?ava=1
258	259	https://sun1-25.userapi.com/impf/uac25CZtNVMkgzoPBCKqUkJZxLEshf3kMcjJFQ/wkPPQF2iEL8.jpg?size=200x0&quality=90&crop=0,0,640,640&sign=22f8dde7d864585ed4243cd62d7c8305&ava=1
259	260	https://sun1-85.userapi.com/impf/c851420/v851420600/92ff1/AfDw-Z72kdE.jpg?size=200x0&quality=90&crop=0,0,503,503&sign=fad8d1ea52af84de8404a556581452cc&ava=1
260	261	https://sun1-24.userapi.com/impf/c622725/v622725010/1e9d6/B8-4P_Y8TGg.jpg?size=200x0&quality=90&crop=67,67,1650,1913&sign=3017fb48ccb745f564e09365e10bee90&ava=1
261	262	https://sun1-94.userapi.com/impf/c637525/v637525011/4bf17/C99-WWuSjrQ.jpg?size=200x0&quality=90&crop=31,31,657,897&sign=340ab71396ef818019c653763b8459c7&ava=1
262	263	https://sun9-31.userapi.com/c10418/u556012/a_1eda7852.jpg?ava=1
263	264	https://sun1-88.userapi.com/impf/c636927/v636927015/17042/4uUTgD2n-e4.jpg?size=200x0&quality=90&crop=265,43,690,690&sign=4573ddd59be62d39fc78fd7434fac60b&ava=1
264	265	https://sun9-8.userapi.com/c10306/u556016/a_d249ca45.jpg?ava=1
265	266	https://sun1-24.userapi.com/impf/QgK_Nj2T_ZN1eoCK4aSZwfVAyplCpNIzTMuK9g/HRibI41Lou0.jpg?size=200x0&quality=90&crop=324,0,1912,1912&sign=cf0833842b2b8a36d3f588a96e6f3a46&ava=1
266	267	https://sun1-47.userapi.com/impf/c850120/v850120489/3125/pM_r7ubqFwo.jpg?size=200x0&quality=90&crop=18,0,492,500&sign=719040ae3363b0354801d44df394a57b&ava=1
267	268	https://sun1-96.userapi.com/impf/c625622/v625622026/1a29/4URJP_xh6Os.jpg?size=200x0&quality=90&crop=0,0,1536,2048&sign=6c9eb52ecf779d59f6853cb9f115b658&ava=1
268	269	https://sun9-40.userapi.com/c4321/u556028/a_2a66af1c.jpg?ava=1
269	270	https://sun1-93.userapi.com/impf/c639226/v639226172/4e5af/4lxcNJ6P1bs.jpg?size=200x0&quality=90&crop=95,1000,941,941&sign=0a5b206dc08284700f7497591f42aa87&ava=1
270	271	https://sun9-23.userapi.com/c09/u556031/a_9de3387.jpg?ava=1
271	272	https://sun1-30.userapi.com/impf/9culMWGUi_FPiSUHAjbC4qhgDknUflMCPZzEWA/RpolfZ0WGvw.jpg?size=200x0&quality=90&crop=0,0,402,402&sign=2ed5c23d13625b07bc65656e840d9150&ava=1
272	273	https://sun1-22.userapi.com/impf/c622521/v622521038/33eaf/OidLAmC-UzE.jpg?size=200x0&quality=90&crop=0,20,200,211&sign=899877368d8128d4303e326d2f2ff899&ava=1
273	274	https://sun9-38.userapi.com/c1424/u556044/a_0528ce9e.jpg?ava=1
274	275	https://sun9-64.userapi.com/c1214/u556045/a_74697408.jpg?ava=1
275	276	https://sun1-25.userapi.com/impf/c851332/v851332610/d41ad/fNipbn-40QA.jpg?size=200x0&quality=90&crop=71,71,1629,2017&sign=521326dba60a3372f35817fb21b965a1&ava=1
276	277	https://sun9-52.userapi.com/c9816/u556050/a_97f3db26.jpg?ava=1
277	278	https://sun1-93.userapi.com/impg/dL3j3t7KL9Af2lYBJ6FKq9IhnaYB2i9WO1jy-A/4SbOVsyIXNY.jpg?size=200x0&quality=90&crop=2,2,540,540&sign=5459cd7b254c561780bd9fb4d01f4683&ava=1
278	279	https://sun1-88.userapi.com/impf/c604826/v604826057/cd47/Fg3F0P_-H7k.jpg?size=200x0&quality=90&crop=0,0,2160,2160&sign=68c10db48eae03b86228bd05641e4c10&ava=1
279	280	https://sun1-20.userapi.com/impf/c855532/v855532990/10fb68/BuTAM13-nXw.jpg?size=200x0&quality=90&crop=0,7,1518,1518&sign=81d096439b9697c266953b33f551fae3&ava=1
280	281	https://sun1-29.userapi.com/impg/zQAFXL7EKrLTM7Kjwq_wSCLS6u2qqAq5vVJ0SQ/QVtulQrb_XU.jpg?size=200x0&quality=90&crop=0,1,960,960&sign=a9e630d6e4f2ba5a86a98921fd02e768&ava=1
281	282	https://sun1-99.userapi.com/impf/4B48ZTjQr0O57hdgCMeEDAqzmV-tLLEMAdeBHQ/LNNAlGXdWrM.jpg?size=200x0&quality=90&crop=330,0,1920,1920&sign=98101ab80c28162909281f43ef75d2db&ava=1
282	283	https://sun1-23.userapi.com/impf/c624118/v624118070/36f36/H02wQhVuFqQ.jpg?size=200x0&quality=90&crop=0,0,957,958&sign=2a980e46f55325bb566ae1d57d7e9fc9&ava=1
283	284	https://sun1-25.userapi.com/impf/c858136/v858136790/2bbd6/qu-yTa16x20.jpg?size=200x0&quality=90&crop=4,4,801,801&sign=477f9f961598226ab42712c1246530ee&ava=1
284	285	https://sun1-24.userapi.com/impg/fWjY8ttWS30EO40icw5izWWPkLpuuAJbnbMcFg/EAJZk3oFTt4.jpg?size=200x0&quality=90&crop=1,1,1496,1496&sign=4a4b1316bd2af9d53e1db6138f238bd6&ava=1
285	286	https://sun9-14.userapi.com/c16/u556081/a_b86184e.jpg?ava=1
286	287	https://sun1-21.userapi.com/impf/IbqNBZ5RmihxLCytDxjOy2a3Q6VWRVxex52glw/qLI4u-3XeXA.jpg?size=200x0&quality=90&crop=0,0,1024,1024&sign=debcd2a99c1f96911708830cc0236bcb&ava=1
287	288	https://sun1-27.userapi.com/impg/M156O2XZgC0GZOAgNwjWSkinX1p2B1xefxHJFQ/uoFhm--C2W4.jpg?size=200x0&quality=90&crop=281,1,957,957&sign=e3e2f3f8381d0d3e569833bd17e122d0&ava=1
288	289	https://sun1-16.userapi.com/impg/GH4RkS83PBjNYEsa3kzxDAJfaVzmMibYM6oAYg/08_zFIKNhL4.jpg?size=200x0&quality=90&crop=126,0,1703,1703&sign=c90af50e253ad8acaa85085dafae4391&ava=1
289	290	https://sun1-30.userapi.com/impg/4y4PsuLl7pPS2bsj3lcVaSlsmwMeN0SAu-HXSw/irVFy8MXhOM.jpg?size=200x0&quality=90&crop=0,274,1008,1008&sign=7c32e1cede311a3cfcf8938a9ea0e78e&ava=1
290	291	https://sun1-91.userapi.com/impf/c834203/v834203955/133cfc/KTB9MtH9yfE.jpg?size=200x0&quality=90&crop=287,0,1535,1535&sign=356fa7bbb9aecb127f59f6fa82391899&ava=1
291	292	https://sun1-96.userapi.com/impf/SKwr5w1mWyFewRjGgXBg0YuiAdLmOdNrcJjmYw/V559EKuVz9w.jpg?size=200x0&quality=90&crop=37,0,850,850&sign=099c27f94a852c28deeede4701c84731&ava=1
292	293	https://sun9-22.userapi.com/c265/u556095/a_ea683226.jpg?ava=1
293	294	https://sun1-16.userapi.com/impf/UcGi5FUz1ZL35OmvXVtBNm2TPi0qIVL5Rjh6Hw/PnDldrL1p-U.jpg?size=200x0&quality=90&crop=33,33,957,957&sign=84ed77f60660894c1137360827d230ae&ava=1
294	295	https://sun1-30.userapi.com/impf/c625828/v625828099/49592/E8fitPVac8s.jpg?size=200x0&quality=90&crop=0,182,960,960&sign=a1313c72c75474d63b242277fb7b9ea4&ava=1
295	296	https://sun9-28.userapi.com/c9826/u556100/a_0fe0cce4.jpg?ava=1
296	297	https://sun1-99.userapi.com/impg/j3pFe5ga0Kbaq1F6lJq3Q4xQudFcu_BCkueh4g/mb_Mut_MiTs.jpg?size=200x0&quality=90&crop=99,217,1151,1151&sign=3ed94ea18888ea45a3db2f938b0c5c7e&ava=1
297	298	https://sun1-21.userapi.com/impf/c840423/v840423677/1629d/4Ok8NAmG5mU.jpg?size=200x0&quality=90&crop=55,0,960,960&sign=1b932dc004a7f15b93501f092cd918ee&ava=1
298	299	https://sun1-14.userapi.com/impf/c622724/v622724105/f733/GGRaz_mdDQE.jpg?size=200x0&quality=90&crop=0,0,959,959&sign=05a237a76a806403523a055b6c41b7bc&ava=1
299	300	https://sun1-85.userapi.com/impf/c845421/v845421817/1264ef/bi0SN7WpO9E.jpg?size=200x0&quality=90&crop=39,524,1323,1323&sign=7a8ee5fcad0ab7aa21d19e1a86f032d8&ava=1
300	301	https://sun9-34.userapi.com/c1218/u556110/a_93aac802.jpg?ava=1
301	302	https://sun1-24.userapi.com/impg/c854216/v854216368/1fe621/yR2g9AK8yDg.jpg?size=200x0&quality=90&crop=7,319,1602,1602&sign=a6455872852254ec696e0b3564fe7620&ava=1
302	303	https://sun1-83.userapi.com/impf/8n7M-JhhxtnUl7OQvoTI8c9oaB4BJw0TS8Cycw/0VA0qecaCRY.jpg?size=200x0&quality=90&crop=597,97,1538,1538&sign=1338a484b66d154a4e871556a5d1d933&ava=1
303	304	https://sun1-28.userapi.com/impf/c636531/v636531119/61b37/r8jW0qEH-zQ.jpg?size=200x0&quality=90&crop=166,0,1538,1538&sign=e19e9d861cfc9a8ab8b027021492fd63&ava=1
304	305	https://sun9-68.userapi.com/c15/u556124/a_0afc0b4.jpg?ava=1
305	306	https://sun1-93.userapi.com/impf/c849132/v849132460/6c5c2/-sXsDPUvJqQ.jpg?size=200x0&quality=90&crop=75,355,1435,1435&sign=09efda1500974645a49d6c8a6190b1a2&ava=1
306	307	https://sun1-24.userapi.com/impf/c636630/v636630130/2a161/g6iu6lZLbv8.jpg?size=200x0&quality=90&crop=787,0,1703,1703&sign=6d2e932a85ccedefaafc2ec36e2a393b&ava=1
307	308	https://sun1-92.userapi.com/impf/c824601/v824601296/38ba3/IBZcxmMkK98.jpg?size=200x0&quality=90&crop=458,0,1440,1440&sign=baaaef1ea235e17177460996b9abe4ab&ava=1
308	309	https://sun1-88.userapi.com/impf/Sz6Vif81v6cP0Z2_fT7RA7WstUGSV1D5rEOsHg/l8kUsGNE_dM.jpg?size=200x0&quality=90&crop=31,31,574,897&sign=181c50238a59822a7e5e97ca913657ea&ava=1
309	310	https://sun9-34.userapi.com/c307907/u556150/a_c613db45.jpg?ava=1
310	311	https://sun9-70.userapi.com/c15/u556159/a_c0545b3.jpg?ava=1
311	312	https://sun9-53.userapi.com/c84/u556164/a_ada621d.jpg?ava=1
312	313	https://sun1-96.userapi.com/impf/c308517/v308517172/5977/ii9I5b6vj0w.jpg?size=200x0&quality=90&crop=26,164,545,545&sign=551008eab2b4f1abb5a10cd24eea28a7&ava=1
313	314	https://sun1-20.userapi.com/impf/c625330/v625330188/39bbe/j4qIxSaU2io.jpg?size=200x0&quality=90&crop=0,0,1153,2048&sign=96ab6a62ccdb3da282d7e5f09432ce7d&ava=1
314	315	https://sun1-17.userapi.com/impf/c639227/v639227191/2c817/CqEm2jBbyfY.jpg?size=200x0&quality=90&crop=0,44,725,725&sign=252fa8632a46bf3714ca0bb3f25efdaf&ava=1
315	316	https://sun1-29.userapi.com/impf/XdvELqcr5lVBOYf4kccBRSfnZRDPkKJ-3K9C9A/z4b5KRI4Dgo.jpg?size=200x0&quality=90&crop=46,46,958,1299&sign=a047c9eada603934f4be77e3693fe8b3&ava=1
316	317	https://sun1-14.userapi.com/impf/c626218/v626218193/272/764jJjoqTEA.jpg?size=200x0&quality=90&crop=0,0,563,681&sign=a8fde6604c1516a246293207122450b0&ava=1
317	318	https://sun9-15.userapi.com/c16/u556202/a_3726c06.jpg?ava=1
318	319	https://sun9-50.userapi.com/c641/u556203/a_bf0168ee.jpg?ava=1
319	320	https://sun9-41.userapi.com/c20/u556204/a_bae55a6.jpg?ava=1
320	321	https://sun1-91.userapi.com/impf/c851028/v851028774/13ba8c/DcI-uaIgrnE.jpg?size=200x0&quality=90&crop=0,124,960,960&sign=ddcc59ad9a80f41767aaaed37fb4a110&ava=1
321	322	https://sun1-17.userapi.com/impf/c627629/v627629210/1e3e4/07oc6BRM_QY.jpg?size=200x0&quality=90&crop=0,3,1435,2157&sign=f1e60e219bd8f28bef39845f8e15cef5&ava=1
322	323	https://sun1-29.userapi.com/impf/c841429/v841429474/38f7e/msRHxisUbDc.jpg?size=200x0&quality=90&crop=0,0,1275,1700&sign=2e3c7a8b8d3dab2cd8252006d2d3ef15&ava=1
323	324	https://sun1-15.userapi.com/impf/c638431/v638431857/4fae5/3M2nYv46CfY.jpg?size=200x0&quality=90&crop=0,118,1371,1371&sign=2969fb8defae874697a0ed4bc0f2ee2c&ava=1
324	325	https://sun1-23.userapi.com/impf/ehyEykYyZyZvS5rdL2qeUr7t_Vc0IapitYoJdg/HfyEKlVyJCQ.jpg?size=200x0&quality=90&crop=0,42,400,547&sign=8a386ae4e8104af254065b36479d50ca&ava=1
325	326	https://sun1-83.userapi.com/impg/c854324/v854324735/193a5f/msloV5ANpnM.jpg?size=200x0&quality=90&crop=40,97,769,1153&sign=8d0e8bc2c1c27f8c7b1711b489e4efce&ava=1
326	327	https://sun1-98.userapi.com/impf/c639516/v639516928/37f55/_PVUgEeKkfg.jpg?size=200x0&quality=90&crop=2136,292,424,425&sign=d8b79354132afc1f8960df7030304a02&ava=1
327	328	https://sun1-20.userapi.com/impf/c626230/v626230248/41730/ga0fjA69wOQ.jpg?size=200x0&quality=90&crop=2,116,1618,1618&sign=4fc69e6bb524382a2051722737ac6dd4&ava=1
328	329	https://sun9-21.userapi.com/c9306/u556255/a_859113aa.jpg?ava=1
329	330	https://sun9-56.userapi.com/c15/u556256/a_7bbf718.jpg?ava=1
330	331	https://sun9-72.userapi.com/c5730/u556259/a_f5a629d7.jpg?ava=1
331	332	https://sun9-28.userapi.com/c118/u556265/a_8379def7.jpg?ava=1
332	333	https://sun1-91.userapi.com/impf/l3Hclf4P4VLfnU4QYYYzxSZA0ayOU-gABynjCw/XaWjpUU-Rx0.jpg?size=200x0&quality=90&crop=0,0,358,479&sign=c297ae1327b88ec335d436cf1281c72e&ava=1
333	334	https://sun1-47.userapi.com/impf/c857236/v857236767/15b70/xvfj8Wz_7B8.jpg?size=200x0&quality=90&crop=0,0,1620,2160&sign=649663a32fb29ad28a432a473382b8f0&ava=1
334	335	https://sun1-98.userapi.com/impf/c858320/v858320731/6a468/2m7es0Jf7Os.jpg?size=200x0&quality=90&crop=103,432,1477,1506&sign=3021c346af38e8fd53a9c482e31b6dec&ava=1
335	336	https://sun1-93.userapi.com/impf/c851028/v851028075/1c4c6b/Q4x91TJwAKY.jpg?size=200x0&quality=90&crop=0,0,840,1256&sign=52a148af01acb6f63ab5e26c18003c8b&ava=1
336	337	https://sun1-26.userapi.com/impf/c845216/v845216313/1d8e79/zSgl2hCGyrQ.jpg?size=200x0&quality=90&crop=193,59,457,457&sign=04613c509936b6ff6930bfb3291cdfbe&ava=1
337	338	https://sun1-92.userapi.com/impf/c636520/v636520287/3c0d/IY8_Ke8IiKI.jpg?size=200x0&quality=90&crop=699,0,1531,1532&sign=32d24cf33c81df2e97c23959f8cbec7d&ava=1
338	339	https://sun1-90.userapi.com/impg/tUj4PRdenKRwg0W4H2C44wufgwb0-17AQDGyHw/m4452gg75_M.jpg?size=200x0&quality=90&crop=3,158,771,771&sign=6f007d8e0a528974887cd3c067fff65b&ava=1
339	340	https://sun1-20.userapi.com/impf/c636619/v636619308/583c6/nclWT6rufjw.jpg?size=200x0&quality=90&crop=280,1,719,719&sign=d522ecf7f4b9e21819a4d156ffe7478a&ava=1
340	341	https://sun1-97.userapi.com/impf/c624727/v624727312/4cf5b/oNOL_gL-anM.jpg?size=200x0&quality=90&crop=0,0,960,960&sign=4f2d6d00c2d29f087fc8a10274e9d73d&ava=1
341	342	https://sun9-71.userapi.com/c592/u556314/a_b3c30f28.jpg?ava=1
342	343	https://sun1-18.userapi.com/impf/smpEcDONSgLxVKyY8ESBtFaAJaBvnllN2oKAOQ/85Sw4ChxybE.jpg?size=200x0&quality=90&crop=203,33,616,616&sign=92a34f395ae0096cfdee41f6ba1a715d&ava=1
343	344	https://sun1-16.userapi.com/impf/c9470/v9470322/2302/ZtrDB22VSEw.jpg?size=200x0&quality=90&crop=127,26,545,545&sign=694e3261c15c57f56ad9700f6317d131&ava=1
344	345	https://sun1-22.userapi.com/impf/c852032/v852032023/16267c/sL6Dc3mJaYU.jpg?size=200x0&quality=90&crop=0,0,1612,2160&sign=45fb5924b170b5b7ef99c9d711730c54&ava=1
345	346	https://sun1-90.userapi.com/impf/c636317/v636317324/5ba8/CDTuC0QrqJ8.jpg?size=200x0&quality=90&crop=0,0,1620,2160&sign=591729e26f35ee9e9aecc3bc9a142b9f&ava=1
346	347	https://sun9-27.userapi.com/c10627/u556327/a_fab7f92c.jpg?ava=1
347	348	https://sun1-84.userapi.com/impf/wlrIIxmdNCSUyWHN0APJDLtCKNzrjzJfUXpu9Q/HFvhVXYAB_4.jpg?size=200x0&quality=90&crop=0,0,960,1280&sign=439d6034e5ea1981390b95306228c3e4&ava=1
348	349	https://sun1-83.userapi.com/impg/Co0zSIrpJenxw4NeaD6QYvUZ50lHU_pqGWx87w/LPaVFKgrF4A.jpg?size=200x0&quality=90&crop=0,27,1242,1242&sign=3984971c3ed0680e79b26d2f54c1697e&ava=1
349	350	https://sun1-85.userapi.com/impg/c854528/v854528475/203e30/ln0ArpOkAm8.jpg?size=200x0&quality=90&crop=0,0,2160,2160&sign=7117fdfc7fcaa25e29dba6f9e931e83b&ava=1
350	351	https://sun9-16.userapi.com/c9760/u556339/a_16f5e9b8.jpg?ava=1
351	352	https://sun9-74.userapi.com/c221/u556340/a_17ceef55.jpg?ava=1
352	353	https://sun1-47.userapi.com/impf/c855636/v855636149/cfc09/oUk8OZ-xk7E.jpg?size=200x0&quality=90&crop=1,1,1071,1071&sign=2a55d8179eeac6c03480b582f9c778a5&ava=1
353	354	https://sun1-28.userapi.com/impg/q2zGEKYpm2SDUNEsP8R8BHHI7-HYBm_u8zUZDA/gAutQwLA3BQ.jpg?size=200x0&quality=90&crop=0,5,1620,2155&sign=4eb7397112384454afdc069b6f2871ed&ava=1
354	355	https://sun1-15.userapi.com/impf/c855632/v855632445/fad85/2zZFhVsdG-s.jpg?size=200x0&quality=90&crop=320,0,1917,1917&sign=1451f1b50f7fcedbec60394e7721008e&ava=1
355	356	https://sun9-17.userapi.com/c4324/u556365/a_0320a772.jpg?ava=1
356	357	https://sun1-24.userapi.com/impg/c855620/v855620706/203664/FYjYvNaKRj4.jpg?size=200x0&quality=90&crop=0,204,1620,1620&sign=0db077d2510ad5a48fac4dce56a649be&ava=1
357	358	https://sun1-91.userapi.com/impf/c840729/v840729544/7fa5d/JF0HcM451T0.jpg?size=200x0&quality=90&crop=0,0,1717,2160&sign=ad7e9d2e9e7ed7d8c724bb8d94b13179&ava=1
358	359	https://sun1-28.userapi.com/impg/gNxlgfhYbf876MWYv0CzMVV6z3fouZAWvJ-dUg/kjeK6eCRWBQ.jpg?size=200x0&quality=90&crop=235,0,958,958&sign=1f5755490ff68563d94a5ecf0cb63150&ava=1
359	360	https://sun1-99.userapi.com/impf/c858416/v858416051/ae33d/L7NChL2Vu2w.jpg?size=200x0&quality=90&crop=0,134,540,540&sign=934cb1fd1eac5c7511812338fabb35c2&ava=1
360	361	https://sun1-19.userapi.com/impf/c844216/v844216293/1c179d/mIgU4PpUp9I.jpg?size=200x0&quality=90&crop=213,0,601,601&sign=8d33f27c1ca55cf815a572e70266eee0&ava=1
361	362	https://sun1-92.userapi.com/impf/c626818/v626818385/2029a/QdbrFXi3ymA.jpg?size=200x0&quality=90&crop=830,245,1462,1462&sign=11dca74bd49f500e5fe7dc263060bfed&ava=1
362	363	https://sun9-22.userapi.com/c168/u556386/a_cab0e36d.jpg?ava=1
363	364	https://sun1-18.userapi.com/impg/GRPBceDKEnonReqaku1-5YbExYrDBa92rBYsJQ/bIdkbJUEdPY.jpg?size=200x0&quality=90&crop=23,23,513,666&sign=b923d1fb3c8d04697d58ae1b36ad54dd&ava=1
364	365	https://sun1-17.userapi.com/impg/jvUPiAoSUnmcHno7UDp08kahwyniBXyRagEAwA/2iM_XF8S-bw.jpg?size=200x0&quality=90&crop=289,4,1915,1915&sign=de4336fb19ee1504da28ccdff08df232&ava=1
365	366	https://sun1-88.userapi.com/impf/c851128/v851128956/992bf/wMovxnLejX4.jpg?size=200x0&quality=90&crop=33,33,684,934&sign=34afc2a995d00b0170f8b697e8df110b&ava=1
366	367	https://sun1-96.userapi.com/impf/g3OiR2-Vog6yFGuk7a4gSgKSxx99l11y248xgg/4tKlS2S42SA.jpg?size=200x0&quality=90&crop=101,25,564,564&sign=cf6058de1f33651ed09887aac66e7606&ava=1
367	368	https://sun9-69.userapi.com/c11/u556394/a_dff7c3b.jpg?ava=1
368	369	https://sun9-18.userapi.com/c16/u556395/a_3091798.jpg?ava=1
369	370	https://sun1-92.userapi.com/impf/c853620/v853620823/170157/w_2gvEwQ23g.jpg?size=200x0&quality=90&crop=0,267,1617,1617&sign=87f6592b71717607c8e136e8be4e1724&ava=1
370	371	https://sun1-85.userapi.com/impf/c857528/v857528365/19afb/baYW4Ljo1R8.jpg?size=200x0&quality=90&crop=78,143,1477,2017&sign=8e7eff16f6be9bce95b5ddf5a4c35deb&ava=1
371	372	https://sun1-15.userapi.com/impf/c845418/v845418556/20eef1/2C_DANv48TA.jpg?size=200x0&quality=90&crop=15,10,2133,2133&sign=4ab43fccf3b6581f8e5c5981ec666270&ava=1
372	373	https://sun1-47.userapi.com/impf/c841225/v841225712/76705/PFcGrZ_uhCM.jpg?size=200x0&quality=90&crop=147,0,585,585&sign=f60b6df228ad8d44884545ee7e6c3673&ava=1
373	374	https://sun1-27.userapi.com/impf/c622519/v622519410/114be/ulF6FJOGDRY.jpg?size=200x0&quality=90&crop=342,461,563,777&sign=adb11868673780439f5abfe0846cdbcb&ava=1
374	375	https://sun9-13.userapi.com/c21/u556414/a_e54efda.jpg?ava=1
375	376	https://sun1-26.userapi.com/impg/ZIKMPjIO41intlaZvH9a00FPVyqYyResz18tSg/hLGOX-1UIYM.jpg?size=200x0&quality=90&crop=0,0,1019,1199&sign=1fb7634a3c50e2a0e951d257cda098a3&ava=1
376	377	https://sun1-86.userapi.com/impf/c840722/v840722640/254dc/ELtM2vaERxM.jpg?size=200x0&quality=90&crop=0,134,384,384&sign=e2db919866a6278b9725a6f2291a8168&ava=1
377	378	https://sun9-40.userapi.com/c9474/u556422/a_2ed1f3b0.jpg?ava=1
378	379	https://sun1-21.userapi.com/impg/qYrKY5lZ1Y-TJaJP9aaPQruloSgLod4woUFE9w/-ukQgJV7Lzs.jpg?size=200x0&quality=90&crop=0,0,319,319&sign=24d7c67c133d974aca252e9c0048ae05&ava=1
379	380	https://sun9-14.userapi.com/c16/u556428/a_75d4b82.jpg?ava=1
380	381	https://sun9-71.userapi.com/c27/u556429/a_578aa9f.jpg?ava=1
381	382	https://sun9-30.userapi.com/c305410/u556430/a_2e65d745.jpg?ava=1
382	383	https://sun1-14.userapi.com/impf/c639224/v639224861/44f25/tTEfun2BsTo.jpg?size=200x0&quality=90&crop=0,0,1536,1536&sign=e3d5cc025f8fca502a1c45dc990ad447&ava=1
383	384	https://sun9-17.userapi.com/c5716/u556441/a_48b059f7.jpg?ava=1
384	385	https://sun1-97.userapi.com/impg/c858232/v858232774/185a92/0-lJoc_277c.jpg?size=200x0&quality=90&crop=20,195,647,647&sign=71a60506cfc8081b3771e8cb22bd7fc4&ava=1
385	386	https://sun1-19.userapi.com/impf/lGW5HjqvvPVWRPi7fhL_Jd8st4e1e8TMB951-A/kMy2n99pkBg.jpg?size=200x0&quality=90&crop=103,0,412,412&sign=cc10ce59cfa50fd50fec076ed2c78ddb&ava=1
386	387	https://sun1-87.userapi.com/impf/c850236/v850236770/1a932b/rcN9kLiO5cE.jpg?size=200x0&quality=90&crop=101,1,877,877&sign=f99b0b556e5080bafdaaaca54071ea32&ava=1
387	388	https://sun1-17.userapi.com/impg/c855216/v855216824/1824a6/b1gz8cn8LS4.jpg?size=200x0&quality=90&crop=0,179,621,621&sign=d2b1c4fcb218affeeccdd99c14102e99&ava=1
388	389	https://sun9-33.userapi.com/c402431/u556461/a_f86609de.jpg?ava=1
389	390	https://sun1-28.userapi.com/impg/R9j77oQuzO2CXVTd8DYZTzUMMo2aE9faFwf9Vw/U77jeajodg0.jpg?size=200x0&quality=90&crop=63,63,1793,1793&sign=c5a02e26158226201a17edc1bfcf61d6&ava=1
390	391	https://sun1-24.userapi.com/impg/c857228/v857228837/d4550/rKqh_XHx31o.jpg?size=200x0&quality=90&crop=0,0,1442,2160&sign=8915216b4fe69624588b12c7dfca40e7&ava=1
391	392	https://sun1-90.userapi.com/impf/c853524/v853524824/8b22e/GBK2GI9YCyo.jpg?size=200x0&quality=90&crop=0,160,959,959&sign=b1421770b05e5920f0da89854f5122cd&ava=1
392	393	https://sun1-18.userapi.com/impf/c633531/v633531480/2ac88/rY0XL5BTE2Y.jpg?size=200x0&quality=90&crop=0,100,600,600&sign=86b0332320ccb18e2ff9efa719873602&ava=1
393	394	https://sun1-98.userapi.com/impf/c845120/v845120263/10cd2d/xNQAn6HE84w.jpg?size=200x0&quality=90&crop=295,68,638,932&sign=7ba805a82b1ea4aed50df93e798ba0f6&ava=1
394	395	https://sun9-23.userapi.com/c14/u556486/a_ef9bc12.jpg?ava=1
395	396	https://sun1-87.userapi.com/impf/c854328/v854328616/8ec8f/t4UHgCHOtVU.jpg?size=200x0&quality=90&crop=71,196,1276,1910&sign=7dd2b029c03ad25df4bc6441efefa8b3&ava=1
396	397	https://sun9-58.userapi.com/c19/u556494/a_0c8840b.jpg?ava=1
397	398	https://sun1-20.userapi.com/impf/c636523/v636523495/4baf7/1sdC2bDpxLQ.jpg?size=200x0&quality=90&crop=256,0,1536,1536&sign=aa28ff96ee56359366ed84993b4d2763&ava=1
398	399	https://sun9-76.userapi.com/c383/u556496/a_bc7f26cd.jpg?ava=1
399	400	https://sun9-52.userapi.com/c08/u556497/a_2f97ded.jpg?ava=1
400	401	https://sun1-86.userapi.com/impf/c604321/v604321501/c2a7/3OH7jorE_L8.jpg?size=200x0&quality=90&crop=423,32,873,874&sign=4ab2980d398e29c86f2f49d52a258c9e&ava=1
401	402	https://sun9-38.userapi.com/c27/u556504/a_15f2725.jpg?ava=1
402	403	https://sun1-27.userapi.com/impg/Hru6Xj3W1XWOMOdB8yXH7jCUg1YhnKFXTjb97Q/u-XJSAqcHaw.jpg?size=200x0&quality=90&crop=644,105,1514,1602&sign=257271ffc51837187d6a6cd2bd279a97&ava=1
403	404	https://sun1-23.userapi.com/impg/c850608/v850608366/ed807/sWT6lnOAylU.jpg?size=200x0&quality=90&crop=0,0,1620,2160&sign=4631ac42494b0a85f6409716a429b32f&ava=1
404	405	https://sun9-32.userapi.com/c1869/u556510/a_1b4bfb52.jpg?ava=1
405	406	https://sun1-87.userapi.com/impf/b6uT1WPCnGVB9Z8OxV_SNb2RIppHiwkkqDTCmQ/hkXVVK5X2kk.jpg?size=200x0&quality=90&crop=0,0,1623,2160&sign=64943e21ef39b01b43246225ff8b6999&ava=1
406	407	https://sun1-95.userapi.com/impf/c850420/v850420627/1aa119/oKCTE9cefL0.jpg?size=200x0&quality=90&crop=2,422,1603,1603&sign=58220d52cac75d6feab6b080b49e7aa9&ava=1
407	408	https://sun9-36.userapi.com/c830/u556522/a_76439e0b.jpg?ava=1
408	409	https://sun1-20.userapi.com/BgxWl201j83sd5yl_bSuOOFH3ndxdmNJaMxfcw/UDaJeANhsuI.jpg?ava=1
409	410	https://sun1-26.userapi.com/impf/TWeq58sigp-BTH333rhvCj7f_4D_WsHK27iv3Q/jfDXQWbX8NU.jpg?size=200x0&quality=90&crop=538,834,857,857&sign=647585c9866efb0a1e08ea83c4a6e1ea&ava=1
410	411	https://sun1-89.userapi.com/impf/c848736/v848736170/13d511/aZF7In8J_OI.jpg?size=200x0&quality=90&crop=77,0,402,402&sign=34831f401d5c3398efc621a0ad9e5b7b&ava=1
411	412	https://sun9-24.userapi.com/c459/u556532/a_15d35e0f.jpg?ava=1
413	414	https://sun1-97.userapi.com/impf/c855416/v855416806/513dc/7wfAV41_qDE.jpg?size=200x0&quality=90&crop=0,201,1536,1536&sign=914f915657d6137cb186f58872cc2a1f&ava=1
414	415	https://sun1-14.userapi.com/impf/c622416/v622416548/ca79/t3XeE6v-Xdk.jpg?size=200x0&quality=90&crop=33,33,701,957&sign=ef37845d6a277797fd397d0303a919f6&ava=1
415	416	https://sun9-37.userapi.com/c08/u556551/a_7f6c114.jpg?ava=1
416	417	https://sun1-18.userapi.com/impf/c622425/v622425556/11013/56EQmBbUaiU.jpg?size=200x0&quality=90&crop=0,0,600,894&sign=b68c532955f94caa63a926626e0eab45&ava=1
417	418	https://sun1-27.userapi.com/impf/c633821/v633821558/31353/IwlG1gPG-qc.jpg?size=200x0&quality=90&crop=0,0,810,1080&sign=8cbf1fe6ee3480967d35991c18e11988&ava=1
418	419	https://sun1-96.userapi.com/impf/c622731/v622731559/18bb8/NMdqTRSEjbU.jpg?size=200x0&quality=90&crop=201,42,806,918&sign=537ecb4a05f0bdd0fb563d6cae847ee3&ava=1
419	420	https://sun1-97.userapi.com/impf/c857624/v857624626/514c0/gHCiNtD04sc.jpg?size=200x0&quality=90&crop=0,0,1536,2048&sign=f47b304fb02329e39b95dfd833662f2c&ava=1
420	421	https://sun1-24.userapi.com/impf/Xz7wSZYJoUC_Jmw3bKt_IgTo_vnTXig-GPxgIg/_ptfgw_zd8k.jpg?size=200x0&quality=90&crop=0,0,960,960&sign=e90ef85bfb400004d1ec313f9ef4e44e&ava=1
421	422	https://sun1-95.userapi.com/impf/c836522/v836522568/318da/uek8XDBUUEU.jpg?size=200x0&quality=90&crop=140,253,882,882&sign=a0a2abfd14c386027eebaed2ff8336a3&ava=1
422	423	https://sun1-94.userapi.com/impf/c637523/v637523569/55e7c/Nw0WUCXTU2o.jpg?size=200x0&quality=90&crop=0,39,623,623&sign=ad5f16a67842df1d083c6ffb89f7fb98&ava=1
423	424	https://sun1-29.userapi.com/impf/c622126/v622126577/16099/KYFZczBeTTc.jpg?size=200x0&quality=90&crop=160,0,960,960&sign=49aa7cf7623f7e103c2cd4baafb06c10&ava=1
424	425	https://sun9-38.userapi.com/c5685/u556579/a_81435f23.jpg?ava=1
425	426	https://sun1-21.userapi.com/impf/c845216/v845216907/be018/oKVojfyTpQs.jpg?size=200x0&quality=90&crop=0,159,960,961&sign=b383241726d10dc7a69c7a66908e2f04&ava=1
426	427	https://sun9-30.userapi.com/c1493/u556590/a_7d577781.jpg?ava=1
427	428	https://sun9-52.userapi.com/c10014/u556593/a_cc4625e3.jpg?ava=1
428	429	https://sun1-18.userapi.com/impf/c836421/v836421599/1663f/AmIRYNyskAU.jpg?size=200x0&quality=90&crop=0,0,1053,1167&sign=f8b5b9ca2690748a5685cd45502b5ab9&ava=1
429	430	https://sun9-58.userapi.com/c10810/u556600/a_734736c9.jpg?ava=1
430	431	https://sun1-83.userapi.com/impf/c849320/v849320567/1a94aa/eHkuyfbku8M.jpg?size=200x0&quality=90&crop=40,311,510,510&sign=c75ab62eadf0566b70ff8aef223c9afe&ava=1
431	432	https://sun9-21.userapi.com/c09/u556603/a_316ff0e.jpg?ava=1
432	433	https://sun1-28.userapi.com/impf/8nlz8ha2zIiOE4vxam1xU8HxSZISkf56rke73A/FiCKg37-wJI.jpg?size=200x0&quality=90&crop=0,0,402,604&sign=ee77b2ef60bc48bdd95bb62068242b9e&ava=1
433	434	https://sun1-19.userapi.com/impf/c638717/v638717618/47193/_7ZQnmjINl0.jpg?size=200x0&quality=90&crop=25,158,390,390&sign=c3d00a39cd21ae3802a03715e923b32f&ava=1
434	435	https://sun1-47.userapi.com/impg/ZH6i5wjvAbuooqMt7vq2eNoBkcgtmP0931PYDA/9gH8yrQPYEk.jpg?size=200x0&quality=90&crop=121,82,1155,1728&sign=b9543fd91344ea3e0b1fc37025507142&ava=1
435	436	https://sun1-26.userapi.com/impf/c622120/v622120624/426ab/P--owO8wxR4.jpg?size=200x0&quality=90&crop=0,0,1200,1200&sign=1220fbe6506760e517832b8f85b9cd51&ava=1
436	437	https://sun9-46.userapi.com/c10972/u556625/a_4632c5bd.jpg?ava=1
437	438	https://sun9-9.userapi.com/c106/u556626/a_5cd3d772.jpg?ava=1
438	439	https://sun1-22.userapi.com/impf/c840530/v840530005/116f4/a3K55906wWM.jpg?size=200x0&quality=90&crop=0,281,627,649&sign=6658d3a78cde247f2c156aff83e9de20&ava=1
439	440	https://sun9-75.userapi.com/c10321/u556631/a_82c2bcc7.jpg?ava=1
440	441	https://sun1-24.userapi.com/impf/c824700/v824700216/7aeda/FLq3-cLRz5o.jpg?size=200x0&quality=90&crop=39,201,650,650&sign=0b6a95199c277035ccbdcecedbfa4f67&ava=1
441	442	https://sun1-30.userapi.com/impf/c633827/v633827637/2e82d/ZAyilIsvGUc.jpg?size=200x0&quality=90&crop=50,0,480,480&sign=b4bdd8c3d085ff20620037b63feb10ed&ava=1
442	443	https://sun1-19.userapi.com/impg/cbi4Jm7C_ttkJlwZCGveqhVqWXorRfQV5FLtAg/j0RcjUj8s9c.jpg?size=200x0&quality=90&crop=94,0,601,601&sign=cd7f2ce5a55d874d8528d94c23c81616&ava=1
443	444	https://sun1-24.userapi.com/impf/c636225/v636225640/2cf61/IIzM1-0rgu8.jpg?size=200x0&quality=90&crop=0,0,791,924&sign=b55e04bda4305e6de318a7e20c9a2c53&ava=1
444	445	https://sun1-47.userapi.com/impg/c853424/v853424362/1acc0f/cKyQTAdKWEE.jpg?size=200x0&quality=90&crop=0,254,1538,1538&sign=4615a2d47fbb6869a7c67240d0f27d70&ava=1
445	446	https://sun1-90.userapi.com/impf/c837125/v837125642/2b50c/J9zLCEBk1rA.jpg?size=200x0&quality=90&crop=406,84,1742,1743&sign=9579a09be11db1fdff6d2cc8e98dc06a&ava=1
446	447	https://sun1-98.userapi.com/impf/c836523/v836523643/33d05/Lloy9xGMZ7U.jpg?size=200x0&quality=90&crop=96,160,1298,1946&sign=310c57049ff90ba31e5ee31e8c71aca0&ava=1
447	448	https://sun1-16.userapi.com/impf/c629307/v629307646/242ef/8REyJfv-ebA.jpg?size=200x0&quality=90&crop=481,0,1624,1624&sign=494ca62a2a8e4b3083f754e1dcdc9f8d&ava=1
448	449	https://sun9-34.userapi.com/c5266/u556647/a_c79f8d44.jpg?ava=1
449	450	https://sun1-91.userapi.com/impf/snyIW4nCYB29T2YRAdqewoiY9a_79sOi6XvOjA/stlgOFMGZII.jpg?size=200x0&quality=90&crop=165,131,371,406&sign=e771cbeca51ed96c059bfe97ad4fff55&ava=1
450	451	https://sun9-75.userapi.com/c409/u556649/a_7318d9bd.jpg?ava=1
451	452	https://sun1-90.userapi.com/impf/c834401/v834401915/7fe84/Ml_MluMMi0o.jpg?size=200x0&quality=90&crop=17,19,970,1288&sign=5f7ba3342f0d89ed21cf2ae85096c9be&ava=1
452	453	https://sun9-38.userapi.com/c11104/u556652/a_753f695d.jpg?ava=1
453	454	https://sun1-97.userapi.com/impf/c623930/v623930653/56f1/-lxoewtS7Aw.jpg?size=200x0&quality=90&crop=0,0,612,612&sign=63a4924d0cdcb4481a0e0a80641178a7&ava=1
454	455	https://sun1-14.userapi.com/impf/HO3dDsfCXiQ1mZR3wC8ZEyf1TKDPqonww4Vndg/1NImLLD_9Lc.jpg?size=200x0&quality=90&crop=0,0,1280,1280&sign=16ef3dd4a20ab0521b3abe24f09aa88d&ava=1
455	456	https://sun1-85.userapi.com/impg/LD-Mdk-iHn0_vt_RQsHWw4Jtt8-5F0WKQ-RTzg/uLIvYAQ5glY.jpg?size=200x0&quality=90&crop=368,76,1576,1579&sign=7419e1fcc171e3712152a3e5292429f5&ava=1
456	457	https://sun1-94.userapi.com/impf/c631322/v631322656/38747/vLFCIOxjXXY.jpg?size=200x0&quality=90&crop=0,0,641,1179&sign=4b11b253e11063eb091867f7d47194f4&ava=1
457	458	https://sun1-18.userapi.com/impf/c858236/v858236747/8ee05/EhpDfr44AYM.jpg?size=200x0&quality=90&crop=71,211,605,605&sign=ab8de40a675633523114841296f31a84&ava=1
458	459	https://sun1-84.userapi.com/impf/c622926/v622926659/5d39b/6dHJc93cHhg.jpg?size=200x0&quality=90&crop=372,67,1538,1538&sign=0fcc15885e3ac438557b60d1c32a4229&ava=1
459	460	https://sun1-24.userapi.com/impf/I7spjV2eR4SVZgauCsYJjdKXhEskEls-r0wBxw/wGnvO2mGgg0.jpg?size=200x0&quality=90&crop=89,0,388,388&sign=96df5dc1a0f66a2fcdf4bcd397ed46aa&ava=1
460	461	https://sun9-9.userapi.com/c4227/u556664/a_118c43d1.jpg?ava=1
461	462	https://sun9-38.userapi.com/c15/u556675/a_f534ce0.jpg?ava=1
462	463	https://sun1-88.userapi.com/impf/AL2qyqYys2mvg1LWJyvLObux2Bsds0oMd0SLRg/jLHkRayz_Zs.jpg?size=200x0&quality=90&crop=19,0,767,1255&sign=6ebc453b50056afe61d76318d2796fcd&ava=1
463	464	https://sun1-87.userapi.com/impg/YaEnUx8JgMembrx9gXg2cCbEusIiPLVX1OqDbQ/bGEOMMCoiqU.jpg?size=200x0&quality=90&crop=1,721,1438,1438&sign=a5b2b96ee3925b2bcc3573b8f7d13471&ava=1
464	465	https://sun1-21.userapi.com/impf/c638527/v638527683/30f9a/PDMOFId_n1M.jpg?size=200x0&quality=90&crop=254,0,960,960&sign=c391049d40884d9756c43a33bf13ac02&ava=1
465	466	https://sun9-46.userapi.com/c14/u556689/a_8f0d13b.jpg?ava=1
466	467	https://sun9-40.userapi.com/c9521/u556701/a_9c8610f5.jpg?ava=1
467	468	https://sun9-8.userapi.com/c16/u556703/a_5b0f02b.jpg?ava=1
468	469	https://sun1-96.userapi.com/impf/c629313/v629313706/1c632/T_m2Qy53j50.jpg?size=200x0&quality=90&crop=21,21,598,598&sign=e0975ea625166a3aa13b0d40baf464b3&ava=1
469	470	https://sun1-96.userapi.com/impf/YmNiqQohucBPvkpEOo8cqbQvfiqiOYo50z1l4A/dgEFRu0mhYc.jpg?size=200x0&quality=90&crop=141,0,797,797&sign=075d74ae7cf03977c6b8f69dea4082b1&ava=1
470	471	https://sun1-91.userapi.com/impf/c844216/v844216255/b9254/MpYUV88A1eE.jpg?size=200x0&quality=90&crop=0,0,1365,2048&sign=75f480293e0bdcf49ca49d72cbaa0658&ava=1
471	472	https://sun9-55.userapi.com/c9771/u556712/a_5342d816.jpg?ava=1
472	473	https://sun1-15.userapi.com/impf/c837337/v837337862/5c9a6/_DsI5Nh1X1Q.jpg?size=200x0&quality=90&crop=198,1,719,719&sign=a01af2cbae83537305a8307cb25fdb6c&ava=1
473	474	https://sun1-93.userapi.com/impg/fGodwb4kLM0NgZSZkit3EyJ1h1QXrEKs6RJ84g/7msSf_pzVrI.jpg?size=200x0&quality=90&crop=0,89,1580,1846&sign=de632a6e39870fb6fecb0afec42fd1df&ava=1
474	475	https://sun1-29.userapi.com/impf/XjMYZQ0q05bv7xFDrnQ1QIaipgl7UZniN7utkA/d-mK2Oog9jA.jpg?size=200x0&quality=90&crop=0,0,1080,1920&sign=ce2b389c250f7a34fc5e96e213b0c0e8&ava=1
475	476	https://sun9-75.userapi.com/c09/u556719/a_162739d.jpg?ava=1
476	477	https://sun1-16.userapi.com/impf/c855736/v855736728/12745e/C3OzNSfKqp0.jpg?size=200x0&quality=90&crop=3,3,633,633&sign=cb926d790729f8974c4a3d0956278dd4&ava=1
477	478	https://sun1-17.userapi.com/impf/c857524/v857524070/45283/ojbDs39L3_c.jpg?size=200x0&quality=90&crop=4,155,949,949&sign=9b17ee18ea2addfa2ee5f14291ff8b54&ava=1
478	479	https://sun1-93.userapi.com/impf/c857720/v857720167/1d934/L8z-jzKGixM.jpg?size=200x0&quality=90&crop=1,44,944,944&sign=d320ea56b556569c633d6eb26c14c8d2&ava=1
479	480	https://sun1-93.userapi.com/impf/c639521/v639521897/5b974/Q1fVVDhb0Kw.jpg?size=200x0&quality=90&crop=241,0,853,853&sign=5fd7111a466e0e3234c507cc82b7b53f&ava=1
480	481	https://sun1-18.userapi.com/impf/nYP4nM9YaGvekNtlsB-mOSRbhvjcYv9JJ7DK0w/8l4xcwTlymg.jpg?size=200x0&quality=90&crop=67,67,1230,1913&sign=496e5ff46d057f4cf12012d71638d47e&ava=1
481	482	https://sun9-23.userapi.com/c538/u556734/a_28fa8e07.jpg?ava=1
482	483	https://sun1-14.userapi.com/impf/c851428/v851428799/1db940/80nczbc3a10.jpg?size=200x0&quality=90&crop=0,6,1104,1104&sign=fe31fe5e4269c4e6f825d26e9e6de7dc&ava=1
483	484	https://sun1-91.userapi.com/impf/c858128/v858128633/416e1/HB8aJNsdtXM.jpg?size=200x0&quality=90&crop=6,6,1265,1265&sign=2837471c037701c138d680c5f814b64f&ava=1
484	485	https://sun1-47.userapi.com/impf/c845418/v845418740/1d77e7/a8Dzya9dKck.jpg?size=200x0&quality=90&crop=0,29,402,402&sign=f846fcec0d48a302820652b1045842d2&ava=1
485	486	https://sun1-99.userapi.com/impf/c855736/v855736284/75751/ViEPaCHgab8.jpg?size=200x0&quality=90&crop=45,45,949,1295&sign=0842535e4a9781ce101739c63129ab6c&ava=1
486	487	https://sun9-6.userapi.com/c18/u556745/a_559f147.jpg?ava=1
487	488	https://sun1-18.userapi.com/impg/b_Yu7xWhsOuZBNAV32PSlsD7GLSWHwcqkcUIyw/5jUd-oZ6alg.jpg?size=200x0&quality=90&crop=49,49,1401,1401&sign=fd1f8687fd48e2ed9887e49eae4ce5bf&ava=1
488	489	https://sun9-43.userapi.com/c128/u556753/a_229fd7a2.jpg?ava=1
489	490	https://sun9-65.userapi.com/c1025/u556754/a_f689d741.jpg?ava=1
490	491	https://sun1-22.userapi.com/impf/ZjCHP9uzZIfZCeHqi4xmz7DQ190_u5ZWKZOsVw/gP6xRyYLkw4.jpg?size=200x0&quality=90&crop=90,0,713,720&sign=c0130f9185458b09cbd3898ee632c13c&ava=1
491	492	https://sun1-47.userapi.com/impf/IMJ388XTUU-tYpl8OoDOXeV2o52eAdBUZ9-i-A/rQKo2WUyRXM.jpg?size=200x0&quality=90&crop=0,0,402,604&sign=0b17ca3f2ca8947c3d6ba9345bd7995e&ava=1
492	493	https://sun9-25.userapi.com/c285/u556761/a_d251e64c.jpg?ava=1
493	494	https://sun1-47.userapi.com/impg/ky7BNOYZiJacBr6ohPHPIBFfZYbxMWZen5oPBA/wlrT0f6sKec.jpg?size=200x0&quality=90&crop=128,2,1916,1916&sign=271e3c819516c88a698440ab9e17f145&ava=1
494	495	https://sun1-17.userapi.com/impg/qlFV_eO5NkXWM8SisnOntnp43WiYtuPA9dGYtA/aUYueU6SdDA.jpg?size=200x0&quality=90&crop=0,0,499,499&sign=6000664aa78726adf54d8504ddc6383e&ava=1
495	496	https://sun9-45.userapi.com/c10516/u556766/a_4af26196.jpg?ava=1
496	497	https://sun1-27.userapi.com/impf/c639117/v639117438/5b9d7/m0oLNw8yvGs.jpg?size=200x0&quality=90&crop=758,0,1802,1920&sign=f61431dc8254e7eef2ff058751475be8&ava=1
497	498	https://sun1-29.userapi.com/impf/c622228/v622228769/1bc1d/SmXXZqJacfY.jpg?size=200x0&quality=90&crop=67,67,1418,1913&sign=dc939f46e6279f12ac44a9b28cb6fdeb&ava=1
498	499	https://sun1-24.userapi.com/impf/iieAkBM2yTST9xSdeCDHFUb4kz474i2F5OD7Wg/e6FUAu8j_10.jpg?size=200x0&quality=90&crop=32,32,610,958&sign=ac388551e0b674cd7fb5f3ae262a88f8&ava=1
499	500	https://sun9-34.userapi.com/c14/u556774/a_b6d0bd5.jpg?ava=1
500	501	https://sun1-97.userapi.com/impf/c851432/v851432670/13c42e/w5T7SNfnQMI.jpg?size=200x0&quality=90&crop=248,1,797,797&sign=bbe57295954b4af83c96668258c5e316&ava=1
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: bsabre
--

COPY public.users (uid, mail, encryptedpass, fname, lname, birth, gender, orientation, bio, avaid, latitude, longitude, interests, status, search_visibility, rating) FROM stdin;
1	admin@gmail.com	54gg9d6	admin	superUser	1989-10-23	male	hetero		\N	\N	\N	{}	not confirmed	f	0
2	user1@gmail.com	fhjegga	Евгений	Жихарев	\N	male	hetero	Евгений  Жихарев Russia Moscow 	1	56.17896	37.545393	{}	confirmed	t	5
3	user2@gmail.com	fhjegga	Павел	Прокудин	\N	male	hetero	Павел  Прокудин Russia Moscow 	2	55.65396	37.345393	{culture,programming,"find something new",youtube}	confirmed	t	5
4	user3@gmail.com	fhjegga	Павел	Кузнецов	1989-01-14	male	homo	Павел  Кузнецов 	3	\N	\N	{}	confirmed	t	7
5	user4@gmail.com	fhjegga	Fenrir	Grey	\N	female	hetero	Fenrir  Grey 	4	\N	\N	{"drink beer",architecture}	confirmed	t	3
6	user5@gmail.com	fhjegga	Алена	Чубаренко	\N	female	homo	Алена  Никитина (Чубаренко) Russia Moscow 	5	55.85396	37.270393	{cooking,architecture}	confirmed	t	11
7	user6@gmail.com	fhjegga	Екатерина	Склифасовская	1984-02-28	female	homo	Екатерина  Склифасовская Russia Moscow 	6	55.32896	37.195393	{}	confirmed	t	0
8	user7@gmail.com	fhjegga	Александр	Дунаев	1987-11-01	male	hetero	Александр  Дунаев Russia Moscow 	7	55.42896	37.445393	{culture,architecture}	confirmed	t	7
9	user8@gmail.com	fhjegga	Александр	Горбачев	\N	male	hetero	Александр  Горбачев Russia Moscow 	8	55.72896	37.245393	{"video games",architecture}	confirmed	t	9
10	user9@gmail.com	fhjegga	Гарик	Салфетков	1987-01-01	male	hetero	Гарик  Салфетков Russia Moscow 	9	55.52896	37.895393	{politics,programming,football}	confirmed	t	10
11	user10@gmail.com	fhjegga	Мартин	Волков	1984-03-02	male		Мартин  Волков Russia Moscow 	10	56.12896	38.120393	{}	confirmed	t	9
12	user11@gmail.com	fhjegga	Иван	Польгов	\N	male	hetero	Иван  Польгов Russia Moscow 	11	55.52896	37.870393	{}	confirmed	t	2
13	user12@gmail.com	fhjegga	Анна	Истошина	\N	female	hetero	Анна  Истошина 	12	\N	\N	{"drink beer",programming,politics}	confirmed	t	5
14	user13@gmail.com	fhjegga	Элина	Саркисян	1985-01-07		hetero	Элина  Саркисян 	13	\N	\N	{}	confirmed	t	7
15	user14@gmail.com	fhjegga	Юрий	Михальченков	1979-05-26	male	hetero	Юрий  Михальченков 	14	\N	\N	{politics,"drink beer"}	confirmed	t	11
16	user15@gmail.com	fhjegga	Дмитрий	Запруднов	1990-09-01	male	hetero	Дмитрий  Запруднов Russia Moscow 	15	56.17896	37.945393	{youtube}	confirmed	t	0
17	user16@gmail.com	fhjegga	Алексей	Новоселов	\N	male	hetero	Алексей  Новоселов 	16	\N	\N	{youtube,"drink beer",cooking}	confirmed	t	9
18	user17@gmail.com	fhjegga	Анатолий	Вересов	\N	male	hetero	Анатолий  Вересов Russia Moscow 	17	55.72896	37.270393	{cooking,architecture}	confirmed	t	7
19	user18@gmail.com	fhjegga	Cadas	Csda	\N	male		Cadas  Csda Russia Moscow 	18	55.77896	37.145393	{"video games",architecture}	confirmed	t	5
20	user19@gmail.com	fhjegga	Елена	Коликова	1985-11-23	female	hetero	Елена  Коликова Russia Moscow 	19	55.47896	37.245393	{cooking}	confirmed	t	4
21	user20@gmail.com	fhjegga	Евгений	Рязанцев	1988-05-17	male	homo	Евгений  Рязанцев 	20	\N	\N	{}	confirmed	t	1
22	user21@gmail.com	fhjegga	Анжелла	Филипенко	1986-09-10	male	hetero	Анжелла  Филипенко Russia Moscow 	21	55.65396	37.220393	{architecture,culture}	confirmed	t	10
23	user22@gmail.com	fhjegga	Лида	Колаева	1987-02-23	female	hetero	Лида  Леснякова (Колаева) Russia Moscow 	22	55.85396	37.520393	{}	confirmed	t	10
24	user23@gmail.com	fhjegga	Мария	Великая	\N	female		Мария  Великая Russia Moscow 	23	55.45396	37.870393	{politics,programming,architecture,football}	confirmed	t	7
25	user24@gmail.com	fhjegga	Роман	Спиридонов	\N	male	hetero	Роман Олегович Спиридонов Russia Moscow 	24	55.85396	37.370393	{cooking}	confirmed	t	3
26	user25@gmail.com	fhjegga	Аня	Воробейкина	\N	female	hetero	Аня  Воробейкина Russia Moscow 	25	55.82896	37.795393	{culture,programming}	confirmed	t	3
27	user26@gmail.com	fhjegga	Саня	Пушкин	1988-05-04	male	hetero	Саня  Пушкин 	26	\N	\N	{football,"drink beer"}	confirmed	t	4
28	user27@gmail.com	fhjegga	Bainnanza	Zzzz	\N	female	hetero	Bainnanza  Zzzz Russia Moscow 	27	55.95396	37.670393	{"video games",culture,architecture}	confirmed	t	10
29	user28@gmail.com	fhjegga	Антон	Тучинский	1984-03-21	male	hetero	Антон  Тучинский Russia Moscow 	28	56.05396	38.095393	{football,cooking}	confirmed	t	2
30	user29@gmail.com	fhjegga	Денис	Скрябин	1988-10-27	male	hetero	Денис  Скрябин Russia Moscow 	29	56.22896	37.420393	{architecture,"find something new",football}	confirmed	t	7
31	user30@gmail.com	fhjegga	Вячеслав	Ткаченко	\N	male	hetero	Вячеслав  Ткаченко 	30	\N	\N	{football,youtube,"video games"}	confirmed	t	3
32	user31@gmail.com	fhjegga	Алина	Пупкина	\N	female	hetero	Алина  Пупкина Russia Moscow 	31	55.77896	37.770393	{"find something new"}	confirmed	t	2
33	user32@gmail.com	fhjegga	Галина	Морозова	1991-02-18	female	homo	Галина  Морозова Russia Moscow 	32	56.07896	37.595393	{"video games",programming,culture}	confirmed	t	10
34	user33@gmail.com	fhjegga	Mary	Sex	1987-05-29	male	hetero	Mary  Sex Russia Moscow 	33	55.52896	37.720393	{}	confirmed	t	8
35	user34@gmail.com	fhjegga	Erik	Akhmetgaliev	\N	female	hetero	Erik  Akhmetgaliev Russia Moscow 	34	55.30396	38.120393	{youtube,football}	confirmed	t	2
36	user35@gmail.com	fhjegga	Сергей	Соколов	1984-02-08	male	hetero	Сергей  Соколов Russia Moscow 	35	55.45396	37.945393	{}	confirmed	t	6
37	user36@gmail.com	fhjegga	Игорь	Михайлов	\N	male	hetero	Игорь Кадет Михайлов 	36	\N	\N	{cooking,football}	confirmed	t	6
38	user37@gmail.com	fhjegga	Анна	Круцак	1986-05-22	female	hetero	Анна  Круцак Russia Moscow 	37	55.45396	37.620393	{cooking,culture,politics}	confirmed	t	11
39	user38@gmail.com	fhjegga	Николай	Жуков	1983-09-06	male	hetero	Николай  Жуков 	38	\N	\N	{architecture}	confirmed	t	1
40	user39@gmail.com	fhjegga	Katja	Patsey	1965-06-05	female	hetero	Katja  Hopmann (Patsey) Russia Moscow 	39	55.95396	38.070393	{culture,"drink beer"}	confirmed	t	7
41	user40@gmail.com	fhjegga	Дмитрий	Фахрутдинов	1978-01-01	male	homo	Дмитрий  Фахрутдинов Russia Moscow 	40	55.82896	37.120393	{architecture}	confirmed	t	1
42	user41@gmail.com	fhjegga	West	Hooligan	1989-04-07	male	hetero	West  Hooligan Russia Moscow 	41	55.27896	37.595393	{programming,cooking,politics}	confirmed	t	1
43	user42@gmail.com	fhjegga	Елена	Гобби	\N	female		Елена  Гобби Russia Moscow 	42	56.22896	37.570393	{"find something new",politics}	confirmed	t	1
44	user43@gmail.com	fhjegga	Наталья	Новикова	1985-03-10	female	hetero	Наталья  Сысоева (Новикова) Russia Moscow 	43	55.32896	37.370393	{cooking}	confirmed	t	7
45	user44@gmail.com	fhjegga	Игорь	Волков	\N	male	homo	Игорь  Волков Russia Moscow 	44	55.67896	37.695393	{programming,"video games"}	confirmed	t	8
46	user45@gmail.com	fhjegga	Telman	Ultovsky	\N		homo	Telman  Ultovsky 	45	\N	\N	{cooking}	confirmed	t	4
47	user46@gmail.com	fhjegga	Женя	Бражник	1987-02-14	male	hetero	Женя  Бражник Russia Moscow 	46	55.55396	37.870393	{"drink beer","find something new",architecture}	confirmed	t	5
48	user47@gmail.com	fhjegga	Кирилл	Спирченко	1982-03-21	male	hetero	Кирилл  Спирченко Russia Moscow 	47	55.65396	37.170393	{culture,football,programming}	confirmed	t	11
49	user48@gmail.com	fhjegga	Дмитрий	Кошелев	1964-06-03	male		Дмитрий  Кошелев Russia Moscow 	48	56.07896	37.770393	{cooking,youtube}	confirmed	t	2
50	user49@gmail.com	fhjegga	Мария	Панина	1983-12-29	female	homo	Мария  Панина Russia Moscow 	49	55.32896	37.695393	{architecture}	confirmed	t	9
51	user50@gmail.com	fhjegga	Анжелика	Петрова	1989-09-13	male	homo	Анжелика  Петрова 	50	\N	\N	{}	confirmed	t	1
52	user51@gmail.com	fhjegga	Виталий	Колобов	\N	male	homo	Виталий  Колобов Russia Moscow 	51	55.70396	37.645393	{"video games",programming}	confirmed	t	4
53	user52@gmail.com	fhjegga	Екатерина	Бражникова	1988-03-05	female	hetero	Екатерина lightsinner Бражникова Russia Moscow 	52	55.35396	37.370393	{"find something new"}	confirmed	t	6
54	user53@gmail.com	fhjegga	Никита	Антипов	\N	male	hetero	Никита  Антипов 	53	\N	\N	{cooking,culture,youtube}	confirmed	t	9
55	user54@gmail.com	fhjegga	Ирина	Егорова	\N	female	hetero	Ирина  Егорова 	54	\N	\N	{"drink beer",youtube}	confirmed	t	7
56	user55@gmail.com	fhjegga	Серега	Рыженков	1988-09-24	male	hetero	Серега  Рыженков Russia Moscow 	55	56.17896	37.570393	{culture}	confirmed	t	7
57	user56@gmail.com	fhjegga	Марина	Шатунова	1987-02-27	female	homo	Марина  Степанова (Шатунова) Russia Moscow 	56	56.12896	37.570393	{programming,politics,"drink beer",architecture}	confirmed	t	4
58	user57@gmail.com	fhjegga	Феликс	Юхтмахер	1985-03-27		hetero	Феликс  Юхтмахер 	57	\N	\N	{}	confirmed	t	8
59	user58@gmail.com	fhjegga	Юля	Муратова	\N	female	hetero	Юля  Муратова 	58	\N	\N	{politics,culture,architecture,cooking}	confirmed	t	5
60	user59@gmail.com	fhjegga	Никита	Макушин	1985-01-07	male	hetero	Никита  Макушин 	59	\N	\N	{}	confirmed	t	3
61	user60@gmail.com	fhjegga	Екатерина	Козлова	\N	female	hetero	Екатерина  Козлова Russia Moscow 	60	55.90396	37.820393	{"find something new",football,"drink beer",culture}	confirmed	t	1
62	user61@gmail.com	fhjegga	Константин	Петров	1983-02-18	male	hetero	Константин  Петров Russia Moscow 	61	55.45396	37.845393	{youtube,"drink beer","find something new"}	confirmed	t	10
63	user62@gmail.com	fhjegga	Виктор	Ильющенко	\N	male	hetero	Виктор  Ильющенко Russia Moscow 	62	55.32896	37.445393	{football,programming,"video games"}	confirmed	t	4
64	user63@gmail.com	fhjegga	Элиза	Лорик	\N	female	hetero	Элиза  Лорик Russia Moscow 	63	56.15396	37.595393	{}	confirmed	t	6
65	user64@gmail.com	fhjegga	Анастасия	Софронова	\N	female	hetero	Анастасия  Софронова Russia Moscow 	64	56.25396	37.345393	{football,"drink beer",architecture}	confirmed	t	0
66	user65@gmail.com	fhjegga	Даша	Алина	\N	female	hetero	Даша  Алина 	65	\N	\N	{"find something new","video games","drink beer",youtube}	confirmed	t	10
67	user66@gmail.com	fhjegga	Максим	Корнеев	1982-09-07	male	hetero	Максим  Корнеев 	66	\N	\N	{cooking,"video games"}	confirmed	t	10
68	user67@gmail.com	fhjegga	Дарья	Исакова	\N	female	homo	Дарья  Исакова Russia Moscow 	67	55.55396	37.245393	{youtube,architecture}	confirmed	t	10
69	user68@gmail.com	fhjegga	Лилия	Давыдова	\N	female	hetero	Лилия  Давыдова Russia Moscow 	68	55.32896	37.745393	{culture,"drink beer"}	confirmed	t	6
70	user69@gmail.com	fhjegga	Daria	Grigoryeva	1991-02-22	male	hetero	Daria  Grigoryeva Russia Moscow 	69	56.17896	37.695393	{"find something new",youtube,programming,football}	confirmed	t	9
71	user70@gmail.com	fhjegga	Андрей	Иванов	\N	male	hetero	Андрей  Иванов Russia Moscow 	70	55.77896	37.345393	{politics,youtube,culture}	confirmed	t	10
72	user71@gmail.com	fhjegga	Павел	Макаров	1987-07-10	male	hetero	Павел  Макаров Russia Moscow 	71	55.35396	37.370393	{"drink beer",politics}	confirmed	t	11
73	user72@gmail.com	fhjegga	Машуля	Амелина	1985-05-30	female	hetero	Машуля  Емельянова (Амелина) 	72	\N	\N	{"video games",youtube,politics}	confirmed	t	8
74	user73@gmail.com	fhjegga	Екатерина	Сиротина	\N	female	hetero	Екатерина  Сиротина Russia Moscow 	73	55.52896	37.295393	{cooking,youtube}	confirmed	t	2
75	user74@gmail.com	fhjegga	Сергей	Дроганов	1979-01-09	male	hetero	Сергей  Дроганов Russia Moscow 	74	56.10396	37.545393	{architecture,politics}	confirmed	t	0
76	user75@gmail.com	fhjegga	Ксюша	Шварц	\N	female	hetero	Ксюша  Шварц Russia Moscow 	75	55.57896	38.095393	{"video games"}	confirmed	t	9
77	user76@gmail.com	fhjegga	Света	Герасименко	1981-02-15	female	hetero	Света  Ферезанова (Герасименко) Russia Moscow 	76	55.80396	38.095393	{}	confirmed	t	5
78	user77@gmail.com	fhjegga	Борис	Васильев	1988-02-04	male		Борис  Васильев Russia Moscow 	77	55.27896	37.770393	{}	confirmed	t	1
79	user78@gmail.com	fhjegga	Александр	Триханов	\N	male		Александр  Триханов 	78	\N	\N	{politics}	confirmed	t	2
80	user79@gmail.com	fhjegga	Ольга	Иванова	\N	female	hetero	Ольга  Макарова (Иванова) 	79	\N	\N	{}	confirmed	t	6
81	user80@gmail.com	fhjegga	Ксения	Дмитриева	\N	female	hetero	Ксения  Дмитриева Russia Moscow 	80	56.05396	38.095393	{politics,architecture,cooking}	confirmed	t	8
82	user81@gmail.com	fhjegga	Алексей	Васильев	1986-11-28	male	homo	Алексей  Васильев 	81	\N	\N	{}	confirmed	t	8
83	user82@gmail.com	fhjegga	Лейла	Иманова	1987-07-27		hetero	Лейла  Иманова Russia Moscow 	82	55.80396	37.445393	{"find something new",cooking,football,architecture}	confirmed	t	4
84	user83@gmail.com	fhjegga	Юля	Пименова	\N	female	homo	Юля  Алферова (Пименова) Russia Moscow 	83	55.67896	37.695393	{architecture}	confirmed	t	4
85	user84@gmail.com	fhjegga	Роман	Батенков	\N	male	hetero	Роман  Батенков Russia Moscow 	84	56.22896	37.995393	{politics,"find something new","drink beer",football}	confirmed	t	1
86	user85@gmail.com	fhjegga	Петр	Аверьянов	1986-10-28	male	hetero	Петр  Аверьянов Russia Moscow 	85	55.82896	37.270393	{"drink beer",politics}	confirmed	t	4
87	user86@gmail.com	fhjegga	Татьяна	Гинь	\N	female	hetero	Татьяна  Гинь Russia Moscow 	86	55.42896	37.120393	{"drink beer","find something new",politics,culture}	confirmed	t	11
88	user87@gmail.com	fhjegga	Виктория	Горенская	1991-05-01	female	homo	Виктория  Горенская Russia Moscow 	87	55.90396	37.820393	{"drink beer",football}	confirmed	t	2
89	user88@gmail.com	fhjegga	Женя	Кузьмина	1990-03-25	female	homo	Женя  Кузьмина Russia Moscow 	88	55.32896	37.420393	{cooking,politics,"drink beer",architecture}	confirmed	t	9
90	user89@gmail.com	fhjegga	Даша	Кускова	1987-09-03	female	homo	Даша Михайловна Кускова Russia Moscow 	89	55.82896	37.620393	{architecture,cooking}	confirmed	t	7
91	user90@gmail.com	fhjegga	Александр	Власенко	1987-09-27	male	hetero	Александр  Власенко 	90	\N	\N	{architecture}	confirmed	t	8
92	user91@gmail.com	fhjegga	Анастасия	Леоненко	1983-06-25	female	homo	Анастасия  Леоненко Russia Moscow 	91	55.82896	37.120393	{culture,architecture}	confirmed	t	10
93	user92@gmail.com	fhjegga	Юрий	Сенин	\N	male	homo	Юрий  Сенин Russia Moscow 	92	56.17896	37.145393	{}	confirmed	t	3
94	user93@gmail.com	fhjegga	Женя	Курмакаев	\N	male	hetero	Женя  Курмакаев Russia Moscow 	93	55.97896	37.520393	{politics}	confirmed	t	2
95	user94@gmail.com	fhjegga	Виталий	Скачков	\N	male	homo	Виталий  Скачков Russia Moscow 	94	55.57896	37.245393	{culture,football}	confirmed	t	5
96	user95@gmail.com	fhjegga	Ев	Гений	1985-10-18	male	hetero	Ев  Гений 	95	\N	\N	{"find something new","video games"}	confirmed	t	0
97	user96@gmail.com	fhjegga	Светлана	Якушенко	1988-02-17	female	hetero	Светлана  Клюева (Якушенко) Russia Moscow 	96	55.72896	37.945393	{culture,football,youtube,programming}	confirmed	t	4
98	user97@gmail.com	fhjegga	Kate	Белоглазова	\N	female	hetero	Kate  Белоглазова Russia Moscow 	97	56.05396	37.895393	{"drink beer",architecture,cooking}	confirmed	t	7
99	user98@gmail.com	fhjegga	Алексей	Красавин	1969-02-26	male	hetero	Алексей KingCrimson Красавин Russia Moscow 	98	55.45396	37.170393	{architecture,"video games","find something new"}	confirmed	t	0
100	user99@gmail.com	fhjegga	Андрей	Суслов	\N	male	hetero	Андрей  Суслов Russia Moscow 	99	55.95396	37.745393	{"drink beer"}	confirmed	t	8
101	user100@gmail.com	fhjegga	Иван	Аникиев	1986-06-26	male	hetero	Иван  Аникиев 	100	\N	\N	{youtube,politics}	confirmed	t	5
102	user101@gmail.com	fhjegga	Вика	Троянова	\N	female	hetero	Вика  Троянова 	101	\N	\N	{}	confirmed	t	6
103	user102@gmail.com	fhjegga	Ирина	Гречишкина	\N	female	hetero	Ирина  Гречишкина Russia Moscow 	102	55.70396	37.170393	{architecture,"video games"}	confirmed	t	5
104	user103@gmail.com	fhjegga	Игорь	Кравен	1990-07-01	male		Игорь  Кравен Russia Moscow 	103	55.80396	37.920393	{"find something new",culture,cooking}	confirmed	t	11
105	user104@gmail.com	fhjegga	Алексей	Кравченко	\N	male	hetero	Алексей  Кравченко Russia Moscow 	104	55.60396	37.270393	{}	confirmed	t	6
106	user105@gmail.com	fhjegga	Денис	Морозов	1982-10-23	male	hetero	Денис  Морозов Russia Moscow 	105	56.15396	37.620393	{youtube,architecture,"drink beer"}	confirmed	t	11
107	user106@gmail.com	fhjegga	Сергей	Ульянов	1982-02-19	male	hetero	Сергей  Ульянов Russia Moscow 	106	55.35396	37.995393	{culture,"video games"}	confirmed	t	8
108	user107@gmail.com	fhjegga	Дмитрий	Болотников	\N	male	homo	Дмитрий  Болотников Russia Moscow 	107	55.35396	37.945393	{culture,programming,youtube}	confirmed	t	11
109	user108@gmail.com	fhjegga	Дарья	Аксенова	1985-09-06	female	homo	Дарья  Золотарева (Аксенова) Russia Moscow 	108	56.25396	37.720393	{culture,football}	confirmed	t	5
110	user109@gmail.com	fhjegga	Иван	Волков	1992-04-06	male	homo	Иван  Волков Russia Moscow 	109	55.80396	37.320393	{}	confirmed	t	5
111	user110@gmail.com	fhjegga	Юлия	Чекунаева	1982-02-25	female	homo	Юлия  Чекунаева Russia Moscow 	110	55.77896	37.870393	{"drink beer",programming}	confirmed	t	10
112	user111@gmail.com	fhjegga	Jonas	Sarakauskas	1984-12-28		hetero	Jonas  Sarakauskas Russia Moscow 	111	56.15396	37.720393	{"find something new",programming}	confirmed	t	6
113	user112@gmail.com	fhjegga	Эка	Абуселидзе	\N	female	hetero	Эка  Абуселидзе Russia Moscow 	112	55.40396	37.245393	{football}	confirmed	t	6
114	user113@gmail.com	fhjegga	Снеговик	Пердовик	\N		hetero	Снеговик  Пердовик Russia Moscow 	113	55.47896	37.920393	{"drink beer",architecture}	confirmed	t	1
115	user114@gmail.com	fhjegga	Roxana	Dojdik	\N		hetero	Roxana  Dojdik Russia Moscow 	114	55.72896	37.595393	{"find something new",politics}	confirmed	t	8
116	user115@gmail.com	fhjegga	Евгений	Василенко	1987-12-21	male	homo	Евгений  Василенко Russia Moscow 	115	56.22896	37.420393	{youtube,football,cooking}	confirmed	t	5
117	user116@gmail.com	fhjegga	Катя	Сарапова	1988-04-24	female	hetero	Катя  Сарапова 	116	\N	\N	{culture,politics,"video games"}	confirmed	t	1
118	user117@gmail.com	fhjegga	Сергей	Бараев	1990-04-23	male	hetero	Сергей  Бараев Russia Moscow 	117	55.97896	37.995393	{culture}	confirmed	t	7
119	user118@gmail.com	fhjegga	Катюша	Нагуманова	\N	female	hetero	Катюша  Нагуманова Russia Moscow 	118	55.62896	37.745393	{"drink beer"}	confirmed	t	6
120	user119@gmail.com	fhjegga	Таня	Сазонова	1981-10-16	female	hetero	Таня  Чуешова (Сазонова) Russia Moscow 	119	56.02896	37.695393	{architecture,football}	confirmed	t	0
121	user120@gmail.com	fhjegga	Vladislav	Nikolaeff	\N	female		Vladislav  Nikolaeff 	120	\N	\N	{youtube,culture,"drink beer"}	confirmed	t	7
122	user121@gmail.com	fhjegga	Антоха	Чащин	\N	male	hetero	Антоха  Чащин 	121	\N	\N	{"find something new",youtube}	confirmed	t	2
123	user122@gmail.com	fhjegga	Сергей	Иванов	\N	male	hetero	Сергей  Иванов Russia Moscow 	122	55.50396	37.470393	{youtube,cooking}	confirmed	t	7
124	user123@gmail.com	fhjegga	Денис	Фахртдинов	1987-06-21	male	hetero	Денис  Фахртдинов Russia Moscow 	123	55.50396	38.120393	{youtube}	confirmed	t	9
125	user124@gmail.com	fhjegga	Динара	Тайтмазова	\N	female	hetero	Динара  Тайтмазова Russia Moscow 	124	55.60396	38.070393	{}	confirmed	t	3
126	user125@gmail.com	fhjegga	Дарья	Лазарева	\N	female	hetero	Дарья  Барканова (Лазарева) 	125	\N	\N	{politics,youtube,cooking}	confirmed	t	1
127	user126@gmail.com	fhjegga	Александра	Сергиенко	\N	male	hetero	Александра  Сергиенко Russia Moscow 	126	55.67896	37.195393	{"find something new","video games",cooking}	confirmed	t	11
128	user127@gmail.com	fhjegga	Violetta	Nazarova	\N		hetero	Violetta  Nazarova Russia Moscow 	127	56.05396	37.445393	{}	confirmed	t	0
129	user128@gmail.com	fhjegga	Игорь	Потаранов	1990-01-27	male	hetero	Игорь  Потаранов Russia Moscow 	128	55.30396	37.220393	{culture}	confirmed	t	1
130	user129@gmail.com	fhjegga	Владимир	Шорохов	\N	male	homo	Владимир  Шорохов 	129	\N	\N	{cooking}	confirmed	t	11
131	user130@gmail.com	fhjegga	Елена	Коробкова	\N	female		Елена  Подгайская (Коробкова) Russia Moscow 	130	55.37896	37.195393	{programming,football,"video games",cooking}	confirmed	t	5
132	user131@gmail.com	fhjegga	Нет	Нет	\N	male	hetero	Нет  Нет Russia Moscow 	131	55.45396	38.020393	{cooking,programming,culture}	confirmed	t	4
133	user132@gmail.com	fhjegga	Bony	Fedia	\N	male	hetero	Bony  Fedia Russia Moscow 	132	55.45396	37.170393	{football}	confirmed	t	5
134	user133@gmail.com	fhjegga	Александра	Ефремова	\N		hetero	Александра  Ефремова Russia Moscow 	133	55.72896	38.095393	{politics,"video games"}	confirmed	t	0
135	user134@gmail.com	fhjegga	Полина	Плахова	1986-08-18	female	hetero	Полина  Мороз (Плахова) Russia Moscow 	134	56.10396	37.145393	{}	confirmed	t	0
136	user135@gmail.com	fhjegga	Евгений	Бобков	1990-09-10	male	hetero	Евгений  Бобков Russia Moscow 	135	56.22896	37.495393	{"video games","find something new"}	confirmed	t	3
137	user136@gmail.com	fhjegga	Жека	Балов	1990-05-28	female		Жека  Балов Russia Moscow 	136	55.40396	37.295393	{architecture,youtube}	confirmed	t	10
138	user137@gmail.com	fhjegga	Константин	Голобоков	1988-07-05	male	homo	Константин  Голобоков Russia Moscow 	137	55.85396	37.895393	{}	confirmed	t	10
139	user138@gmail.com	fhjegga	Гиви	Гиви	1967-05-28	female		Гиви  Гиви Russia Moscow 	138	55.85396	37.745393	{politics,"video games"}	confirmed	t	6
140	user139@gmail.com	fhjegga	Леонид	Сергеевич	1987-09-13	male	hetero	Леонид  Сергеевич Russia Moscow 	139	56.15396	37.820393	{"video games",football,architecture}	confirmed	t	5
141	user140@gmail.com	fhjegga	Александр	Константинов	\N	male	hetero	Александр  Константинов Russia Moscow 	140	55.65396	37.420393	{football,"drink beer",politics}	confirmed	t	0
142	user141@gmail.com	fhjegga	Наталья	Валюшкина	\N	female	homo	Наталья Natalyushkina Валюшкина Russia Moscow 	141	56.12896	37.520393	{}	confirmed	t	10
143	user142@gmail.com	fhjegga	Иван	Корнеев	\N	male	hetero	Иван  Корнеев Russia Moscow 	142	55.75396	37.920393	{politics,architecture,cooking}	confirmed	t	8
144	user143@gmail.com	fhjegga	Victor	Kalashnikov	1958-08-13		hetero	Victor  Kalashnikov Ukraine Kiev 	143	50.3786	30.6914	{youtube,"video games"}	confirmed	t	0
145	user144@gmail.com	fhjegga	Александр	Жуков	1984-07-04	male	hetero	Александр Cat Жуков Russia Moscow 	144	55.80396	37.820393	{}	confirmed	t	2
146	user145@gmail.com	fhjegga	Майя	Зейн	\N	male	homo	Майя  Зейн Russia Moscow 	145	56.10396	37.145393	{programming,cooking,football,"drink beer"}	confirmed	t	9
147	user146@gmail.com	fhjegga	Вера	Каруличева	\N	female	hetero	Вера  Каруличева Russia Moscow 	146	56.10396	37.945393	{programming,"find something new",culture}	confirmed	t	0
148	user147@gmail.com	fhjegga	Никита	Клименко	1992-07-01	male		Никита  Клименко Ukraine Kiev 	147	50.0536	30.0664	{culture}	confirmed	t	10
149	user148@gmail.com	fhjegga	Виталик	Осипенков	\N	male	hetero	Виталик  Осипенков Russia Moscow 	148	55.30396	37.445393	{architecture,"find something new",football}	confirmed	t	6
150	user149@gmail.com	fhjegga	Света	Марголин	1987-06-07	male	hetero	Света  Марголин 	149	\N	\N	{programming,"drink beer"}	confirmed	t	5
151	user150@gmail.com	fhjegga	Анастасия	Кострова	1989-10-02	female	hetero	Анастасия  Веретенникова (Кострова) Russia Moscow 	150	55.52896	37.170393	{"find something new",football}	confirmed	t	3
152	user151@gmail.com	fhjegga	Max	Max	\N		hetero	Max  Max 	151	\N	\N	{youtube,"video games"}	confirmed	t	1
153	user152@gmail.com	fhjegga	Мария	Туникова	\N	female	hetero	Мария  Туникова Russia Moscow 	152	56.17896	37.870393	{}	confirmed	t	2
154	user153@gmail.com	fhjegga	Леся	Смирнова	\N	female	hetero	Леся  Смирнова Russia Moscow 	153	56.10396	37.670393	{politics}	confirmed	t	10
155	user154@gmail.com	fhjegga	Ольга	Шилякова	\N	female	hetero	Ольга  Шилякова Russia Moscow 	154	56.22896	37.595393	{politics,"drink beer","find something new",cooking}	confirmed	t	2
156	user155@gmail.com	fhjegga	Лилия	Занько	1981-02-27	female	hetero	Лилия  Ласкина (Занько) Russia Moscow 	155	56.07896	37.620393	{cooking,youtube,culture,"video games"}	confirmed	t	10
157	user156@gmail.com	fhjegga	Марина	Петрова	\N	female	hetero	Марина  Синишина (Петрова) Russia Moscow 	156	56.10396	38.045393	{architecture,cooking}	confirmed	t	10
158	user157@gmail.com	fhjegga	Анастасия	Мазанова	\N	female	homo	Анастасия  Мазанова 	157	\N	\N	{}	confirmed	t	8
159	user158@gmail.com	fhjegga	Дмитрий	Бучинский	\N	male	hetero	Дмитрий  Бучинский 	158	\N	\N	{"video games",youtube,architecture}	confirmed	t	9
160	user159@gmail.com	fhjegga	Мария	Фрося	\N	female	homo	Мария  Фрося Russia Moscow 	159	55.35396	37.370393	{politics}	confirmed	t	11
161	user160@gmail.com	fhjegga	Антон	Земляков	\N	male	hetero	Антон  Земляков Russia Moscow 	160	55.27896	37.270393	{}	confirmed	t	9
162	user161@gmail.com	fhjegga	Юля	Миронова	\N	female	homo	Юля  Миронова Russia Moscow 	161	55.25396	37.345393	{football,politics}	confirmed	t	10
163	user162@gmail.com	fhjegga	Игорь	Жилинков	\N	male	hetero	Игорь  Жилинков Russia Moscow 	162	55.47896	37.820393	{}	confirmed	t	8
164	user163@gmail.com	fhjegga	Алексей	Пряхин	1983-03-29	male	hetero	Алексей  Пряхин 	163	\N	\N	{politics,youtube,"find something new"}	confirmed	t	4
165	user164@gmail.com	fhjegga	Олег	Беляков	1986-03-13	male	hetero	Олег  Беляков Russia Moscow 	164	55.62896	37.420393	{}	confirmed	t	0
166	user165@gmail.com	fhjegga	Надежда	Петровская	\N	female		Надежда  Петровская Russia Moscow 	165	55.35396	37.520393	{}	confirmed	t	0
167	user166@gmail.com	fhjegga	Юлия	Коченова	1988-05-16	female	hetero	Юлия  Коченова Russia Moscow 	166	55.42896	38.070393	{}	confirmed	t	7
168	user167@gmail.com	fhjegga	Илья	Костецкий	\N	male	hetero	Илья  Костецкий 	167	\N	\N	{culture,politics}	confirmed	t	7
169	user168@gmail.com	fhjegga	Артур	Кабиров	1986-12-23	male	hetero	Артур  Кабиров Russia Moscow 	168	55.57896	37.745393	{}	confirmed	t	7
170	user169@gmail.com	fhjegga	Пустая	Нетов	1987-05-05	male	hetero	Пустая страница Нетов Russia Moscow 	169	56.05396	37.170393	{"find something new",cooking}	confirmed	t	0
171	user170@gmail.com	fhjegga	Сергей	Образцов	1984-09-02	male	hetero	Сергей  Образцов Russia Moscow 	170	55.57896	37.820393	{culture}	confirmed	t	9
172	user171@gmail.com	fhjegga	Katy	Амелина	1989-03-11	female	hetero	Katy  Амелина 	171	\N	\N	{youtube}	confirmed	t	3
173	user172@gmail.com	fhjegga	Елена	Вишневская	1988-10-11	female	hetero	Елена  Вишневская Russia Moscow 	172	56.05396	38.045393	{}	confirmed	t	5
174	user173@gmail.com	fhjegga	Дарья	Дружинина	\N	female	hetero	Дарья  Дружинина Russia Moscow 	173	55.92896	38.070393	{football,politics}	confirmed	t	11
175	user174@gmail.com	fhjegga	Павел	Угольков	1988-04-06	male	hetero	Павел  Угольков Russia Moscow 	174	56.15396	37.295393	{"drink beer"}	confirmed	t	9
176	user175@gmail.com	fhjegga	Михаил	Балашов	1977-05-17	male	hetero	Михаил  Балашов Russia Moscow 	175	55.72896	37.745393	{culture,"video games"}	confirmed	t	3
177	user176@gmail.com	fhjegga	Костя	Пильник	\N	male	hetero	Костя  Пильник Russia Moscow 	176	55.55396	38.045393	{architecture,"find something new",football}	confirmed	t	8
178	user177@gmail.com	fhjegga	Елена	Исхакова	\N	female	hetero	Елена  Сушкова (Исхакова) Russia Moscow 	177	55.62896	37.395393	{architecture,politics,football,youtube}	confirmed	t	10
179	user178@gmail.com	fhjegga	Александр	Приходько	1984-12-07	male		Александр  Приходько Russia Moscow 	178	55.67896	37.970393	{}	confirmed	t	9
180	user179@gmail.com	fhjegga	Илья	Эйхлер	1987-02-16	male	hetero	Илья  Эйхлер Russia Moscow 	179	55.42896	37.245393	{football}	confirmed	t	8
181	user180@gmail.com	fhjegga	Евгений	Зозуля	1989-11-21	male	hetero	Евгений  Зозуля Russia Moscow 	180	55.30396	38.120393	{politics}	confirmed	t	10
182	user181@gmail.com	fhjegga	Виктория	Люминарская	1988-01-14	female	hetero	Виктория  Люминарская Russia Moscow 	181	55.52896	37.395393	{"find something new",cooking}	confirmed	t	9
183	user182@gmail.com	fhjegga	Елизавета	Буйнова	1992-07-10	female	hetero	Елизавета  Буйнова 	182	\N	\N	{football}	confirmed	t	8
184	user183@gmail.com	fhjegga	Ксюшенька	Руднева	1988-06-29			Ксюшенька  Руднева Russia Moscow 	183	55.27896	37.970393	{"drink beer",youtube,programming,football}	confirmed	t	4
185	user184@gmail.com	fhjegga	Алексей	Сочнев	1987-12-16	male	homo	Алексей  Сочнев Russia Moscow 	184	55.70396	37.745393	{programming}	confirmed	t	4
186	user185@gmail.com	fhjegga	Илья	Драмчег	1987-05-24	male	homo	Илья  Драмчег Russia Moscow 	185	55.77896	37.270393	{}	confirmed	t	7
187	user186@gmail.com	fhjegga	Анна	Матюхина	1983-06-22	female	hetero	Анна  Матюхина Russia Moscow 	186	56.22896	37.970393	{"video games","drink beer",culture}	confirmed	t	9
188	user187@gmail.com	fhjegga	Дилярочка	Нурина	\N	female	hetero	Дилярочка  Нурина Russia Moscow 	187	55.72896	37.695393	{}	confirmed	t	1
189	user188@gmail.com	fhjegga	Раш	Калабергено	\N	male	hetero	Раш  Калабергено Russia Moscow 	188	56.10396	37.470393	{"drink beer",football,"find something new"}	confirmed	t	3
190	user189@gmail.com	fhjegga	Александр	Бармин	1987-10-02	male	hetero	Александр  Бармин 	189	\N	\N	{}	confirmed	t	5
191	user190@gmail.com	fhjegga	Ivan	Bakusov	\N	female	hetero	Ivan Alexandrovich Bakusov Russia Moscow 	190	55.27896	37.570393	{}	confirmed	t	4
192	user191@gmail.com	fhjegga	Мария	Федорова	\N	female	hetero	Мария  Федорова Russia Moscow 	191	55.50396	37.895393	{cooking,"drink beer"}	confirmed	t	6
193	user192@gmail.com	fhjegga	Игорь	Храпов	1981-08-29	male	hetero	Игорь  Храпов 	192	\N	\N	{}	confirmed	t	10
194	user193@gmail.com	fhjegga	Александр	Ковязин	1987-12-29	male	hetero	Александр  Ковязин Russia Moscow 	193	55.72896	37.845393	{"drink beer","find something new"}	confirmed	t	8
195	user194@gmail.com	fhjegga	Владимир	Просин	\N	male	homo	Владимир  Просин Russia Moscow 	194	55.72896	38.095393	{"video games",architecture}	confirmed	t	5
196	user195@gmail.com	fhjegga	Леонид	Свирин	\N			Леонид  Свирин 	195	\N	\N	{"drink beer",football,"video games"}	confirmed	t	1
197	user196@gmail.com	fhjegga	Борис	Трушин	1982-12-14	male	hetero	Борис  Трушин 	196	\N	\N	{football}	confirmed	t	10
198	user197@gmail.com	fhjegga	Даниил	Руденко	1989-08-23	male	hetero	Даниил  Руденко Russia Moscow 	197	56.22896	37.270393	{}	confirmed	t	3
199	user198@gmail.com	fhjegga	Roman	Parkhomenko	1990-10-11	female	homo	Roman  Parkhomenko 	198	\N	\N	{youtube}	confirmed	t	6
200	user199@gmail.com	fhjegga	Дмитрий	Юзефович	1990-01-06	male	hetero	Дмитрий  Юзефович Russia Moscow 	199	55.47896	37.870393	{}	confirmed	t	2
201	user200@gmail.com	fhjegga	Елизавета	Кузнецова	\N	female	hetero	Елизавета  Кузнецова Russia Moscow 	200	56.22896	37.395393	{cooking,politics}	confirmed	t	8
202	user201@gmail.com	fhjegga	Мария	Евграфова	1986-03-27	female	hetero	Мария  Евграфова Russia Moscow 	201	55.40396	37.945393	{"drink beer",culture}	confirmed	t	11
203	user202@gmail.com	fhjegga	Татьянка	Смирнова	1988-07-29		hetero	Татьянка  Смирнова Russia Moscow 	202	56.00396	37.720393	{architecture,politics,"video games"}	confirmed	t	3
204	user203@gmail.com	fhjegga	Марк	Холодцов	1991-08-12	female	homo	Марк  Холодцов Russia Moscow 	203	55.52896	37.170393	{}	confirmed	t	0
205	user204@gmail.com	fhjegga	Егор	Голюк	\N	male	homo	Егор  Голюк Russia Moscow 	204	55.85396	37.520393	{"drink beer",cooking}	confirmed	t	8
206	user205@gmail.com	fhjegga	Евгений	Богатиков	1989-08-30	male	hetero	Евгений  Богатиков Russia Moscow 	205	55.80396	37.970393	{"find something new","drink beer",football,culture}	confirmed	t	3
207	user206@gmail.com	fhjegga	Дмитрий	Иванов	1981-05-26	male	homo	Дмитрий  Иванов 	206	\N	\N	{culture,"find something new",programming}	confirmed	t	3
208	user207@gmail.com	fhjegga	Katusha	Goldedition	\N	male	hetero	Katusha  Goldedition Russia Moscow 	207	55.25396	37.670393	{cooking,football}	confirmed	t	0
209	user208@gmail.com	fhjegga	Василий	Ожмегов	1982-02-21	male	hetero	Василий  Ожмегов Russia Moscow 	208	55.95396	37.295393	{politics,youtube,football,culture}	confirmed	t	8
210	user209@gmail.com	fhjegga	Елена	Розанова	1983-08-05	female	homo	Елена  Розанова Russia Moscow 	209	56.25396	38.120393	{}	confirmed	t	1
211	user210@gmail.com	fhjegga	Эдуард	Семенец	\N		hetero	Эдуард  Семенец Russia Moscow 	210	55.95396	37.170393	{football}	confirmed	t	2
212	user211@gmail.com	fhjegga	Андрей	Долгов	1990-03-19	male	hetero	Андрей Fake Долгов Russia Moscow 	211	55.92896	37.845393	{culture,"find something new",politics}	confirmed	t	11
213	user212@gmail.com	fhjegga	Алла	Любчик	1994-02-25	female	homo	Алла  Любчик Russia Moscow 	212	55.35396	37.245393	{}	confirmed	t	9
214	user213@gmail.com	fhjegga	Иван	Степанюк	\N	male	hetero	Иван kompromiss Степанюк Russia Moscow 	213	55.87896	38.120393	{"drink beer",culture,architecture,cooking}	confirmed	t	5
215	user214@gmail.com	fhjegga	Марина	Коробейникова	\N	female	hetero	Марина  Коробейникова Russia Moscow 	214	55.62896	37.395393	{culture,"drink beer"}	confirmed	t	11
216	user215@gmail.com	fhjegga	Таня	Маликова	\N	male	hetero	Таня  Маликова 	215	\N	\N	{cooking,youtube}	confirmed	t	10
217	user216@gmail.com	fhjegga	Татьяна	Лысенко	\N	female	hetero	Татьяна  Лысенко Russia Moscow 	216	55.27896	38.070393	{football}	confirmed	t	5
218	user217@gmail.com	fhjegga	Владимир	Довженко	\N	male	hetero	Владимир  Довженко 	217	\N	\N	{programming,politics}	confirmed	t	8
219	user218@gmail.com	fhjegga	Елена	Саул	\N	female	hetero	Елена  Саул Russia Moscow 	218	55.65396	37.245393	{cooking}	confirmed	t	10
220	user219@gmail.com	fhjegga	Денис	Нахват	\N	male	hetero	Денис  Нахват 	219	\N	\N	{youtube,"find something new"}	confirmed	t	2
221	user220@gmail.com	fhjegga	София	Холодок	1986-01-28	female	hetero	София  Калинина (Холодок) Russia Moscow 	220	56.05396	37.820393	{programming,"find something new",architecture}	confirmed	t	5
222	user221@gmail.com	fhjegga	Инесса	Гаврильчева	1988-03-29	male	hetero	Инесса  Гаврильчева Russia Moscow 	221	55.30396	37.270393	{}	confirmed	t	5
223	user222@gmail.com	fhjegga	Анна	Манукова	1986-09-17	female	hetero	Анна  Манукова Russia Moscow 	222	55.30396	37.295393	{architecture,culture,politics}	confirmed	t	0
224	user223@gmail.com	fhjegga	Веталь	Ole	1983-06-17	male	hetero	Веталь  Ole Russia Moscow 	223	55.47896	37.470393	{programming}	confirmed	t	2
225	user224@gmail.com	fhjegga	Тимур	Урманчеев	1987-04-29	male		Тимур  Урманчеев Russia Moscow 	224	56.17896	37.895393	{"drink beer"}	confirmed	t	10
226	user225@gmail.com	fhjegga	Глафир	Хренников	\N		hetero	Глафир  Хренников 	225	\N	\N	{"find something new"}	confirmed	t	9
227	user226@gmail.com	fhjegga	Ванька	Пашков	1989-10-06	female	hetero	Ванька  Пашков Russia Moscow 	226	55.90396	37.245393	{culture,youtube}	confirmed	t	4
228	user227@gmail.com	fhjegga	Дмитрий	Семенов	1984-11-12	male	hetero	Дмитрий  Семенов Russia Moscow 	227	56.07896	37.820393	{}	confirmed	t	5
229	user228@gmail.com	fhjegga	Anna	Kokina	\N	female	hetero	Anna  Kokina Russia Moscow 	228	55.42896	37.745393	{}	confirmed	t	2
230	user229@gmail.com	fhjegga	Ирина	Кривченок	\N	female	hetero	Ирина  Кривченок Russia Moscow 	229	55.25396	37.645393	{politics,football}	confirmed	t	3
231	user230@gmail.com	fhjegga	Евгений	Крюков	1907-01-01	male	homo	Евгений  Крюков Russia Moscow 	230	55.95396	37.795393	{"video games",cooking,youtube}	confirmed	t	1
232	user231@gmail.com	fhjegga	Екатерина	Томасова	\N	female	hetero	Екатерина  Томасова Russia Moscow 	231	56.12896	37.220393	{politics,programming,"find something new"}	confirmed	t	0
233	user232@gmail.com	fhjegga	Наталья	Коренева	\N	female	hetero	Наталья  Коренева 	232	\N	\N	{"video games"}	confirmed	t	8
234	user233@gmail.com	fhjegga	Михаил	Саблин	1984-12-17	male	hetero	Михаил  Саблин Russia Moscow 	233	55.32896	37.870393	{youtube,culture,"find something new"}	confirmed	t	7
235	user234@gmail.com	fhjegga	Надежда	Страшко	\N	female	hetero	Надежда  Страшко (Страшко) Russia Moscow 	234	55.37896	37.345393	{politics}	confirmed	t	9
236	user235@gmail.com	fhjegga	Сергей	Ухин	1985-03-01	male	homo	Сергей  Ухин Russia Moscow 	235	55.47896	37.395393	{"drink beer"}	confirmed	t	5
237	user236@gmail.com	fhjegga	Максим	Кокорев	\N	male	hetero	Максим  Кокорев 	236	\N	\N	{"video games",cooking,programming,culture}	confirmed	t	6
238	user237@gmail.com	fhjegga	Mari	Mari	\N	female	hetero	Mari  Mari Ukraine Kiev 	237	50.3286	30.241400000000002	{programming}	confirmed	t	9
239	user238@gmail.com	fhjegga	Анастасия	Ахрамеева	1988-06-28	female	hetero	Анастасия  Ахрамеева Russia Moscow 	238	56.10396	37.320393	{youtube,culture,programming}	confirmed	t	7
240	user239@gmail.com	fhjegga	Иван	Смирнов	1983-09-18	male	hetero	Иван  Смирнов Russia Moscow 	239	55.67896	38.045393	{football,"find something new"}	confirmed	t	11
241	user240@gmail.com	fhjegga	Дарья	Загвоздина	1987-01-13	female	hetero	Дарья  Сонькина (Загвоздина) 	240	\N	\N	{culture,politics,football,programming}	confirmed	t	6
242	user241@gmail.com	fhjegga	Ринат	Валиев	1983-10-06	male	hetero	Ринат  Валиев Russia Moscow 	241	56.05396	37.220393	{}	confirmed	t	5
243	user242@gmail.com	fhjegga	Андрей	Иванов	1992-07-25	male	homo	Андрей  Иванов Russia Moscow 	242	55.87896	37.870393	{football,politics}	confirmed	t	11
244	user243@gmail.com	fhjegga	Наташа	Дерябкина	1987-08-23		hetero	Наташа  Дерябкина Russia Moscow 	243	56.00396	38.045393	{football,culture,politics,architecture}	confirmed	t	10
245	user244@gmail.com	fhjegga	Михаил	Шаврак	1987-04-17	male	hetero	Михаил  Шаврак Russia Moscow 	244	55.92896	37.820393	{architecture,youtube,cooking,politics}	confirmed	t	2
246	user245@gmail.com	fhjegga	Максим	Лобачев	1983-12-02	male	hetero	Максим  Лобачев Russia Moscow 	245	55.87896	37.695393	{}	confirmed	t	10
247	user246@gmail.com	fhjegga	Слава	Иванеко	1989-10-12	male	hetero	Слава  Иванеко 	246	\N	\N	{politics,architecture,"video games"}	confirmed	t	8
248	user247@gmail.com	fhjegga	Михаил	Апциаури	1989-04-19	male	hetero	Михаил  Апциаури 	247	\N	\N	{youtube}	confirmed	t	0
249	user248@gmail.com	fhjegga	Саша	Алексин	1986-10-09	male	hetero	Саша  Алексин Russia Moscow 	248	56.17896	37.695393	{politics,"find something new","drink beer"}	confirmed	t	6
250	user249@gmail.com	fhjegga	Мария	Итигина	\N	female	homo	Мария  Поротникова (Итигина) 	249	\N	\N	{politics,youtube,"find something new",culture}	confirmed	t	6
251	user250@gmail.com	fhjegga	Марина	Каверина	\N	female	hetero	Марина  Каверина Russia Moscow 	250	55.80396	37.195393	{cooking}	confirmed	t	7
252	user251@gmail.com	fhjegga	Алексей	Баталин	1977-01-04	male	hetero	Алексей  Баталин Russia Moscow 	251	55.45396	37.170393	{"drink beer",youtube,culture}	confirmed	t	2
253	user252@gmail.com	fhjegga	Алексей	Харченко	1989-07-26	male	hetero	Алексей  Харченко Russia Moscow 	252	56.15396	38.095393	{"find something new",programming}	confirmed	t	8
254	user253@gmail.com	fhjegga	Юлия	Козлова	1986-03-12	female		Юлия  Козлова Russia Moscow 	253	55.40396	37.670393	{"find something new"}	confirmed	t	0
255	user254@gmail.com	fhjegga	Дмитрий	Петров	1987-11-16	male	hetero	Дмитрий  Петров Russia Moscow 	254	56.22896	37.595393	{}	confirmed	t	9
256	user255@gmail.com	fhjegga	Валентина	Литвинова	\N	female	homo	Валентина  Литвинова Russia Moscow 	255	56.22896	37.845393	{culture,"video games",football}	confirmed	t	7
257	user256@gmail.com	fhjegga	Анастасия	Соболева	1989-03-09	female	homo	Анастасия  Азарова (Соболева) 	256	\N	\N	{"drink beer"}	confirmed	t	3
258	user257@gmail.com	fhjegga	Кастяй	Буркин	\N	female	hetero	Кастяй  Буркин Russia Moscow 	257	55.85396	37.520393	{cooking,"find something new"}	confirmed	t	9
259	user258@gmail.com	fhjegga	Патя	Денгаева	1989-10-30	female	hetero	Патя  Денгаева Russia Moscow 	258	55.72896	37.645393	{}	confirmed	t	6
260	user259@gmail.com	fhjegga	Разная	Настоящая	1984-11-05	male	hetero	Разная  Настоящая Russia Moscow 	259	55.80396	37.170393	{youtube,"find something new",architecture}	confirmed	t	9
261	user260@gmail.com	fhjegga	Андрей	Островерхов	\N	male	hetero	Андрей  Островерхов Ukraine Kiev 	260	50.5786	30.5414	{"find something new",politics}	confirmed	t	7
262	user261@gmail.com	fhjegga	Екатерина	Кириллова	\N	female	hetero	Екатерина  Кириллова Russia Moscow 	261	55.25396	37.995393	{programming,"video games"}	confirmed	t	3
263	user262@gmail.com	fhjegga	Артур	Майоров	1990-09-13	male	hetero	Артур  Майоров Russia Moscow 	262	55.25396	37.995393	{}	confirmed	t	1
264	user263@gmail.com	fhjegga	Мария	Анциферова	1986-08-18	female	hetero	Мария  Анциферова 	263	\N	\N	{}	confirmed	t	0
265	user264@gmail.com	fhjegga	Павел	Родюков	1981-07-09	male	hetero	Павел  Родюков Russia Moscow 	264	56.15396	37.870393	{"find something new"}	confirmed	t	11
266	user265@gmail.com	fhjegga	Сергей	Тонков	\N	male	hetero	Сергей  Тонков Russia Moscow 	265	56.15396	37.370393	{cooking,youtube,politics}	confirmed	t	2
267	user266@gmail.com	fhjegga	Rimma	Mustafina	\N		hetero	Rimma  Mustafina 	266	\N	\N	{}	confirmed	t	0
268	user267@gmail.com	fhjegga	Наталья	Андреева	\N	female	hetero	Наталья  Андреева 	267	\N	\N	{architecture,programming}	confirmed	t	0
269	user268@gmail.com	fhjegga	Александр	Шадчнев	\N	male	hetero	Александр  Шадчнев Russia Moscow 	268	55.57896	37.745393	{programming,politics}	confirmed	t	2
270	user269@gmail.com	fhjegga	Анютка	Стрижеченко	1987-10-04	female	hetero	Анютка  Тюткалова (Стрижеченко) Russia Moscow 	269	55.45396	37.645393	{programming,football}	confirmed	t	9
271	user270@gmail.com	fhjegga	Игнат	Прохоров	\N			Игнат  Прохоров Russia Moscow 	270	56.00396	37.145393	{culture,football,"drink beer",architecture}	confirmed	t	8
272	user271@gmail.com	fhjegga	Витя	Сизов	\N	male	homo	Витя  Сизов Russia Moscow 	271	55.72896	38.045393	{}	confirmed	t	7
273	user272@gmail.com	fhjegga	Сергей	Королев	1962-01-31	male	homo	Сергей  Королев Russia Moscow 	272	55.37896	37.295393	{"video games",programming}	confirmed	t	0
274	user273@gmail.com	fhjegga	Виталий	Яблочкин	1987-03-22	male	homo	Виталий  Яблочкин Russia Moscow 	273	56.10396	38.070393	{"video games",youtube}	confirmed	t	9
275	user274@gmail.com	fhjegga	Доминик	Ховард	1977-12-07	male	hetero	Доминик  Ховард Russia Moscow 	274	55.72896	37.120393	{architecture,youtube,"drink beer"}	confirmed	t	5
276	user275@gmail.com	fhjegga	Александр	Андреев	1983-04-04	male	hetero	Александр  Андреев Russia Moscow 	275	55.95396	37.670393	{politics}	confirmed	t	11
277	user276@gmail.com	fhjegga	Жуй	Лю	\N	female	hetero	Жуй  Лю Russia Moscow 	276	56.15396	37.770393	{youtube,politics,culture}	confirmed	t	1
278	user277@gmail.com	fhjegga	Ириха	Ромашева	\N	male	hetero	Ириха  Ромашева Russia Moscow 	277	55.90396	37.295393	{}	confirmed	t	4
279	user278@gmail.com	fhjegga	Полина	Шилова	1988-09-20	female	homo	Полина  Шилова Russia Moscow 	278	55.32896	37.870393	{politics,"find something new",programming}	confirmed	t	0
280	user279@gmail.com	fhjegga	Ольга	Митяшина	\N	female	homo	Ольга  Митяшина Russia Moscow 	279	55.82896	37.520393	{cooking}	confirmed	t	11
281	user280@gmail.com	fhjegga	Ирина	Ефанова	1985-10-06	female	homo	Ирина  Ефанова 	280	\N	\N	{programming}	confirmed	t	9
282	user281@gmail.com	fhjegga	Андрей	Мелащенко	1984-01-25	male	homo	Андрей  Мелащенко Russia Moscow 	281	56.22896	37.270393	{football,"video games","find something new"}	confirmed	t	9
283	user282@gmail.com	fhjegga	Лена	Левицкая	\N	female	homo	Лена  Левицкая Russia Moscow 	282	55.72896	37.695393	{"drink beer"}	confirmed	t	2
284	user283@gmail.com	fhjegga	Дмитрий	Шеленин	\N	male	hetero	Дмитрий  Шеленин Russia Moscow 	283	55.82896	37.770393	{architecture,culture,programming}	confirmed	t	2
285	user284@gmail.com	fhjegga	Мария	Бубенщикова	1985-06-26	female	hetero	Мария  Зайцева (Бубенщикова) Russia Moscow 	284	56.02896	38.120393	{}	confirmed	t	9
286	user285@gmail.com	fhjegga	Апельсин	Заводной	\N	female	hetero	Апельсин  Заводной Russia Moscow 	285	56.20396	37.595393	{politics,"find something new",football,architecture}	confirmed	t	8
287	user286@gmail.com	fhjegga	Георгий	Тимченко	1987-01-15	male	hetero	Георгий  Тимченко Russia Moscow 	286	55.92896	37.145393	{}	confirmed	t	9
288	user287@gmail.com	fhjegga	Виктория	Чернобровкина	\N	female	hetero	Виктория  Леонова (Чернобровкина) Russia Moscow 	287	55.32896	37.470393	{}	confirmed	t	6
289	user288@gmail.com	fhjegga	Дмитрий	Хомич	1983-06-10	male	hetero	Дмитрий  Хомич Russia Moscow 	288	55.85396	38.045393	{programming,cooking,"find something new"}	confirmed	t	2
290	user289@gmail.com	fhjegga	Аксинья	Гусева	1983-08-17	female	hetero	Аксинья  Микеничева (Гусева) Russia Moscow 	289	55.62896	37.970393	{football}	confirmed	t	6
291	user290@gmail.com	fhjegga	Наташа	Варкки	\N	female	hetero	Наташа  Варкки 	290	\N	\N	{football,"drink beer"}	confirmed	t	1
292	user291@gmail.com	fhjegga	Дмитрий	Кравцов	1988-04-14	male	hetero	Дмитрий  Кравцов Russia Moscow 	291	55.37896	37.420393	{football,culture,cooking,"find something new"}	confirmed	t	3
293	user292@gmail.com	fhjegga	Артур	Логинов	1991-05-12	male	hetero	Артур  Логинов Russia Moscow 	292	55.90396	37.870393	{programming,"video games"}	confirmed	t	9
294	user293@gmail.com	fhjegga	Анюта	Браду	\N	female	hetero	Анюта  Калачева (Браду) Russia Moscow 	293	56.05396	37.745393	{"video games",cooking,culture}	confirmed	t	9
295	user294@gmail.com	fhjegga	Марина	Матвеева	\N	female	hetero	Марина  Матвеева Russia Moscow 	294	55.60396	37.295393	{"drink beer"}	confirmed	t	4
296	user295@gmail.com	fhjegga	Ирина	Сапрыкина	\N	female	hetero	Ирина  Сапрыкина Russia Moscow 	295	56.10396	37.470393	{"video games"}	confirmed	t	2
297	user296@gmail.com	fhjegga	Маргарита	Морозова	\N	female	hetero	Маргарита  Сазонова (Морозова) Russia Moscow 	296	55.85396	37.420393	{"find something new",cooking}	confirmed	t	3
298	user297@gmail.com	fhjegga	Ирина	Гуторова	\N	female	hetero	Ирина  Гуторова Russia Moscow 	297	55.80396	37.920393	{football,"find something new",architecture}	confirmed	t	0
299	user298@gmail.com	fhjegga	Илья	Лебедев	\N	male	hetero	Илья  Лебедев 	298	\N	\N	{}	confirmed	t	8
300	user299@gmail.com	fhjegga	Святослав	Панов	1992-01-21	male	hetero	Святослав  Панов Russia Moscow 	299	55.87896	37.270393	{"drink beer",architecture,culture,football}	confirmed	t	8
301	user300@gmail.com	fhjegga	Виктория	Грузнова	1989-07-26	female	hetero	Виктория  Грузнова Russia Moscow 	300	55.82896	37.845393	{"find something new",cooking}	confirmed	t	11
302	user301@gmail.com	fhjegga	Дарья	Процкая	\N	female	hetero	Дарья  Процкая Russia Moscow 	301	55.52896	37.745393	{}	confirmed	t	2
303	user302@gmail.com	fhjegga	Ирина	Арутюнова	1985-11-23	female	hetero	Ирина  Беляева (Арутюнова) Russia Moscow 	302	55.35396	37.120393	{}	confirmed	t	1
304	user303@gmail.com	fhjegga	Татьяна	Владимировна	1989-12-10	female	homo	Татьяна  Владимировна Russia Moscow 	303	55.90396	37.345393	{culture,football}	confirmed	t	10
305	user304@gmail.com	fhjegga	Михаил	Степанов	\N	male	homo	Михаил  Степанов Russia Moscow 	304	55.67896	37.270393	{cooking,youtube,programming}	confirmed	t	9
306	user305@gmail.com	fhjegga	Clockwork	Mary	1988-09-11	female	hetero	Clockwork  Mary Russia Moscow 	305	55.27896	37.945393	{culture,architecture,youtube,programming}	confirmed	t	5
307	user306@gmail.com	fhjegga	Даня	Мельников	\N	male	hetero	Даня  Мельников Russia Moscow 	306	55.87896	37.245393	{"drink beer",politics}	confirmed	t	9
308	user307@gmail.com	fhjegga	Дима	Чепиков	\N	female	homo	Дима  Чепиков Russia Moscow 	307	55.65396	37.470393	{"find something new",youtube}	confirmed	t	11
309	user308@gmail.com	fhjegga	Ирина	Кюнер	1980-07-01	female		Ирина  Кюнер Russia Moscow 	308	55.67896	37.445393	{football,"video games",cooking}	confirmed	t	5
310	user309@gmail.com	fhjegga	Олег	Клейменов	1985-07-01	male	hetero	Олег  Клейменов Russia Moscow 	309	55.95396	37.295393	{"find something new",politics,cooking}	confirmed	t	3
311	user310@gmail.com	fhjegga	Real	Mox	\N	male	hetero	Real  Mox Russia Moscow 	310	55.37896	37.845393	{}	confirmed	t	4
312	user311@gmail.com	fhjegga	Bars	Mishenev	\N	female	homo	Bars  Mishenev 	311	\N	\N	{football}	confirmed	t	8
313	user312@gmail.com	fhjegga	Даша	Бочарова	1988-02-27	female	hetero	Даша  Бочарова Russia Moscow 	312	55.82896	37.345393	{cooking,football}	confirmed	t	4
314	user313@gmail.com	fhjegga	Дмитрий	Ильин	\N	male	hetero	Дмитрий  Ильин 	313	\N	\N	{"video games",politics}	confirmed	t	0
315	user314@gmail.com	fhjegga	Анна	Петрова	1987-11-16	female	hetero	Анна  Петрова Russia Moscow 	314	56.17896	37.470393	{architecture,culture,"drink beer","find something new"}	confirmed	t	11
316	user315@gmail.com	fhjegga	Дмитрий	Сузан	1972-11-26	male		Дмитрий жесть Сузан Russia Moscow 	315	55.45396	37.945393	{politics,cooking}	confirmed	t	6
317	user316@gmail.com	fhjegga	Любовь	Иванова	\N	female	hetero	Любовь Lynx Иванова Russia Moscow 	316	55.27896	37.445393	{}	confirmed	t	6
318	user317@gmail.com	fhjegga	Lolwtfbbq	Lolwtfbbq	\N	male	hetero	Lolwtfbbq  Lolwtfbbq Russia Moscow 	317	56.12896	37.945393	{"find something new",culture}	confirmed	t	5
319	user318@gmail.com	fhjegga	Алексей	Шибанов	1984-10-28	male	hetero	Алексей  Шибанов Russia Moscow 	318	55.57896	38.095393	{}	confirmed	t	10
320	user319@gmail.com	fhjegga	Алена	Китаева	\N	male	hetero	Алена  Китаева Russia Moscow 	319	56.00396	37.620393	{cooking,football,"find something new"}	confirmed	t	4
321	user320@gmail.com	fhjegga	Anna	Lavrentyeva	\N	male		Anna  Lavrentyeva Russia Moscow 	320	56.12896	37.970393	{architecture,culture,"find something new"}	confirmed	t	0
322	user321@gmail.com	fhjegga	Дария	Григоренко	1988-07-02	female	hetero	Дария  Парсаданова (Григоренко) Russia Moscow 	321	55.70396	37.570393	{programming}	confirmed	t	7
323	user322@gmail.com	fhjegga	Галя	Михайлова	\N	female	hetero	Галя  Михайлова 	322	\N	\N	{}	confirmed	t	11
324	user323@gmail.com	fhjegga	Юлия	Shultz	\N	female	homo	Юлия  Shultz Russia Moscow 	323	55.85396	37.120393	{football,"video games"}	confirmed	t	4
325	user324@gmail.com	fhjegga	Магомед	Хачалов	\N		hetero	Магомед  Хачалов Russia Moscow 	324	55.30396	37.820393	{architecture}	confirmed	t	6
326	user325@gmail.com	fhjegga	Felix	Shpilman	1985-10-02	female	hetero	Felix  Shpilman Russia Moscow 	325	55.85396	37.495393	{"find something new",architecture}	confirmed	t	8
327	user326@gmail.com	fhjegga	Валерий	Грузинов	\N	male	hetero	Валерий  Грузинов Russia Moscow 	326	56.15396	37.295393	{}	confirmed	t	4
328	user327@gmail.com	fhjegga	Alyona	Пелешко	\N		homo	Alyona  Пелешко Russia Moscow 	327	55.85396	37.920393	{football}	confirmed	t	1
329	user328@gmail.com	fhjegga	Александр	Громов	\N	male	homo	Александр  Громов 	328	\N	\N	{programming,"find something new"}	confirmed	t	2
330	user329@gmail.com	fhjegga	Колян	Батов	\N	female	hetero	Колян  Батов Russia Moscow 	329	55.32896	37.645393	{}	confirmed	t	0
331	user330@gmail.com	fhjegga	Валерий	Одинцов	\N	male	homo	Валерий Vell Одинцов Russia Moscow 	330	55.42896	37.445393	{}	confirmed	t	10
332	user331@gmail.com	fhjegga	Мишка	Белый	1930-01-21		hetero	Мишка  Белый Russia Moscow 	331	55.90396	37.470393	{}	confirmed	t	2
333	user332@gmail.com	fhjegga	Хусейн	Кубанов	\N	male	hetero	Хусейн  Кубанов Russia Moscow 	332	55.37896	37.920393	{politics,culture,youtube,"drink beer"}	confirmed	t	8
334	user333@gmail.com	fhjegga	Lidi	Smirnova	1990-11-05		hetero	Lidi  Smirnova Russia Moscow 	333	56.12896	37.620393	{cooking,programming,culture,"drink beer"}	confirmed	t	2
335	user334@gmail.com	fhjegga	Юля	Каретина	\N	female	hetero	Юля  Каретина 	334	\N	\N	{architecture,programming}	confirmed	t	3
336	user335@gmail.com	fhjegga	Елена	Колядина	1988-03-13	female	hetero	Елена  Колядина Russia Moscow 	335	55.87896	37.945393	{"drink beer","video games",architecture}	confirmed	t	2
337	user336@gmail.com	fhjegga	Тимофей	Арканников	\N	male	hetero	Тимофей  Арканников 	336	\N	\N	{}	confirmed	t	6
338	user337@gmail.com	fhjegga	Илья	Ермолаев	1988-01-09	male		Илья  Ермолаев Russia Moscow 	337	55.77896	37.495393	{}	confirmed	t	6
339	user338@gmail.com	fhjegga	Ульяна	Стаханова	\N	male	hetero	Ульяна  Стаханова Russia Moscow 	338	55.97896	37.720393	{"video games"}	confirmed	t	0
340	user339@gmail.com	fhjegga	Оксана	Оленина	\N	female	hetero	Оксана счастливая Оленина Russia Moscow 	339	56.15396	38.020393	{youtube,football,programming}	confirmed	t	4
341	user340@gmail.com	fhjegga	Катюша	Букарина	1986-05-27	male	hetero	Катюша  Букарина Russia Moscow 	340	56.02896	37.220393	{}	confirmed	t	9
342	user341@gmail.com	fhjegga	Елена	Климкина	\N	female	homo	Елена  Климкина Russia Moscow 	341	55.95396	37.945393	{"drink beer",youtube,football}	confirmed	t	3
343	user342@gmail.com	fhjegga	Анна	Рейштат	1987-03-21	female	hetero	Анна  Рейштат Russia Moscow 	342	56.12896	37.420393	{cooking}	confirmed	t	7
344	user343@gmail.com	fhjegga	Александр	Цыбиногин	\N	male	homo	Александр  Цыбиногин 	343	\N	\N	{"find something new",youtube}	confirmed	t	9
345	user344@gmail.com	fhjegga	Антон	Головин	\N	male	hetero	Антон Алексеевич Головин Russia Moscow 	344	55.37896	37.945393	{football,programming}	confirmed	t	9
346	user345@gmail.com	fhjegga	Станислав	Браварник	\N	male	hetero	Станислав  Браварник Russia Moscow 	345	56.02896	37.595393	{}	confirmed	t	6
347	user346@gmail.com	fhjegga	Ксения	Фадеева	\N	female	homo	Ксения  Фадеева Russia Moscow 	346	55.42896	37.470393	{politics}	confirmed	t	7
348	user347@gmail.com	fhjegga	Анастасия	Костина	1984-12-28	female	hetero	Анастасия  Костина Russia Moscow 	347	56.07896	37.970393	{architecture}	confirmed	t	11
349	user348@gmail.com	fhjegga	Сергей	Зарезин	1902-01-01	male	hetero	Сергей  Зарезин Russia Moscow 	348	56.12896	37.395393	{culture,politics}	confirmed	t	4
350	user349@gmail.com	fhjegga	Екатерина	Гриценко	\N	female	homo	Екатерина  Гриценко Russia Moscow 	349	56.07896	37.545393	{politics,cooking,"drink beer","video games"}	confirmed	t	5
351	user350@gmail.com	fhjegga	Зенит	Чемпион	\N	male	hetero	Зенит  Чемпион Russia Moscow 	350	55.30396	37.820393	{architecture}	confirmed	t	10
352	user351@gmail.com	fhjegga	Асем	Тулегенова	1989-02-03	female	hetero	Асем  Тулегенова 	351	\N	\N	{}	confirmed	t	10
353	user352@gmail.com	fhjegga	Наталья	Мамаева	\N	female	hetero	Наталья  Мамаева Russia Moscow 	352	56.02896	37.770393	{programming,politics}	confirmed	t	7
354	user353@gmail.com	fhjegga	Александр	Китаев	1983-10-31	male	hetero	Александр  Китаев Russia Moscow 	353	55.65396	37.520393	{}	confirmed	t	11
355	user354@gmail.com	fhjegga	Игорь	Бабич	\N	male	hetero	Игорь  Бабич Russia Moscow 	354	55.47896	37.995393	{architecture,youtube}	confirmed	t	9
356	user355@gmail.com	fhjegga	Мария	Пащенко	\N	female		Мария  Пащенко Russia Moscow 	355	55.65396	37.695393	{cooking,"find something new"}	confirmed	t	0
357	user356@gmail.com	fhjegga	Кирилл	Васильев	1987-05-06	male	hetero	Кирилл  Васильев Russia Moscow 	356	56.25396	37.970393	{youtube,culture,football,"find something new"}	confirmed	t	1
358	user357@gmail.com	fhjegga	Илья	Артеев	\N	male	homo	Илья  Артеев 	357	\N	\N	{football}	confirmed	t	1
359	user358@gmail.com	fhjegga	Ксения	Попова	\N	female	homo	Ксения  Шабанова (Попова) Russia Moscow 	358	56.22896	37.895393	{programming,politics,architecture}	confirmed	t	9
360	user359@gmail.com	fhjegga	Макс	Акопов	\N	male	hetero	Макс  Акопов 	359	\N	\N	{culture}	confirmed	t	0
361	user360@gmail.com	fhjegga	Анна	Юлкина	1987-12-24	female	homo	Анна  Юлкина Russia Moscow 	360	55.42896	37.570393	{}	confirmed	t	3
362	user361@gmail.com	fhjegga	Яна	Буренок	\N	female	hetero	Яна  Буренок Russia Moscow 	361	55.72896	37.495393	{}	confirmed	t	9
363	user362@gmail.com	fhjegga	Леха	Попов	1985-09-06	female	hetero	Леха  Попов Russia Moscow 	362	55.65396	37.445393	{"video games",architecture}	confirmed	t	2
364	user363@gmail.com	fhjegga	Максим	Кузнецов	\N	male	hetero	Максим  Кузнецов Russia Moscow 	363	56.17896	37.295393	{}	confirmed	t	7
365	user364@gmail.com	fhjegga	Иван	Яковлев	1981-07-09	male	homo	Иван Владимирович Яковлев Russia Moscow 	364	55.57896	37.770393	{}	confirmed	t	4
366	user365@gmail.com	fhjegga	Юлия	Мацуева	\N	female	hetero	Юлия  Мацуева 	365	\N	\N	{"drink beer"}	confirmed	t	4
367	user366@gmail.com	fhjegga	Василий	Пиманов	\N	male	homo	Василий  Пиманов 	366	\N	\N	{youtube,"find something new"}	confirmed	t	9
368	user367@gmail.com	fhjegga	Олимпия	Иванова	1980-07-15	female	hetero	Олимпия  Иванова 	367	\N	\N	{"video games"}	confirmed	t	9
369	user368@gmail.com	fhjegga	Потеря	общества	\N	male	hetero	Потеря  Для общества Russia Moscow 	368	55.45396	37.445393	{"find something new",football,culture}	confirmed	t	8
370	user369@gmail.com	fhjegga	Андрей	Иванов	1980-08-30	male	hetero	Андрей  Иванов Russia Moscow 	369	55.75396	37.820393	{}	confirmed	t	8
371	user370@gmail.com	fhjegga	Алексей	Сидоров	1973-07-30	male	homo	Алексей  Сидоров Russia Moscow 	370	56.22896	37.970393	{cooking,architecture,programming}	confirmed	t	10
372	user371@gmail.com	fhjegga	Denis	Genisev	1987-01-27	female	hetero	Denis  Genisev Russia Moscow 	371	55.70396	37.195393	{architecture,"find something new",culture}	confirmed	t	11
373	user372@gmail.com	fhjegga	Алмаз	Якупов	\N	female	homo	Алмаз  Якупов Russia Moscow 	372	55.40396	37.845393	{"find something new",cooking}	confirmed	t	7
374	user373@gmail.com	fhjegga	Ирина	Губина	\N	female	hetero	Ирина  Губина Russia Moscow 	373	55.27896	37.395393	{politics,youtube}	confirmed	t	11
375	user374@gmail.com	fhjegga	Лиза	Ларина	1985-12-13	female	hetero	Лиза  Ларина Russia Moscow 	374	55.90396	38.120393	{"find something new",football}	confirmed	t	5
376	user375@gmail.com	fhjegga	Вика	Белоусова	\N	female	homo	Вика  Белоусова 	375	\N	\N	{programming,cooking}	confirmed	t	10
377	user376@gmail.com	fhjegga	Евгений	Николашичев	1986-11-19	male	hetero	Евгений  Николашичев Russia Moscow 	376	55.70396	37.795393	{"find something new",cooking,"drink beer",architecture}	confirmed	t	9
378	user377@gmail.com	fhjegga	Максим	Андреев	1986-06-03	male	hetero	Максим  Андреев Russia Moscow 	377	55.52896	37.220393	{football,"find something new",culture}	confirmed	t	3
379	user378@gmail.com	fhjegga	Евгений	Осипов	\N	male	hetero	Евгений  Осипов Russia Moscow 	378	55.92896	37.645393	{}	confirmed	t	0
380	user379@gmail.com	fhjegga	Костя	Ефимов	\N	male	hetero	Костя  Ефимов Russia Moscow 	379	55.35396	37.895393	{culture}	confirmed	t	5
381	user380@gmail.com	fhjegga	Денис	Макаров	\N	male	hetero	Денис  Макаров Russia Moscow 	380	55.47896	37.945393	{programming}	confirmed	t	7
382	user381@gmail.com	fhjegga	Александр	Гаврилов	1988-09-14	male	hetero	Александр Sash Гаврилов Russia Moscow 	381	56.15396	37.545393	{youtube,cooking}	confirmed	t	6
383	user382@gmail.com	fhjegga	Олег	Меркурьев	\N	male	hetero	Олег  Меркурьев Russia Moscow 	382	55.32896	37.645393	{"drink beer",architecture}	confirmed	t	3
384	user383@gmail.com	fhjegga	Валерий	Богданов	\N	male		Валерий  Богданов Russia Moscow 	383	55.25396	37.820393	{youtube,culture}	confirmed	t	8
385	user384@gmail.com	fhjegga	Мишаня	Аверкин	1989-06-02	female	hetero	Мишаня  Аверкин Russia Moscow 	384	56.22896	37.370393	{architecture,politics,"find something new"}	confirmed	t	1
386	user385@gmail.com	fhjegga	Александра	Крыжановская	1986-01-14	female	homo	Александра  Крыжановская (Крыжановская) 	385	\N	\N	{architecture,"drink beer"}	confirmed	t	5
387	user386@gmail.com	fhjegga	Юрий	Лавренов	\N	male	homo	Юрий  Лавренов Russia Moscow 	386	55.45396	37.995393	{programming,"drink beer",football}	confirmed	t	0
388	user387@gmail.com	fhjegga	Сергей	Трапезников	1978-12-26	male	hetero	Сергей  Трапезников Russia Moscow 	387	55.52896	38.095393	{"drink beer","video games"}	confirmed	t	5
389	user388@gmail.com	fhjegga	Юрий	Сахаров	\N	male	hetero	Юрий  Сахаров Russia Moscow 	388	55.45396	38.020393	{programming,cooking}	confirmed	t	2
390	user389@gmail.com	fhjegga	Александра	Лебедева	\N	male		Александра  Лебедева Russia Moscow 	389	56.22896	37.295393	{programming}	confirmed	t	5
391	user390@gmail.com	fhjegga	Лола	Жалолова	\N	male	hetero	Лола  Жалолова Russia Moscow 	390	55.47896	37.920393	{}	confirmed	t	2
392	user391@gmail.com	fhjegga	Юлия	Пузикова	1990-02-01	female		Юлия  Прохорова (Пузикова) Russia Moscow 	391	56.25396	37.995393	{}	confirmed	t	1
393	user392@gmail.com	fhjegga	Роза	Шмуклер	1986-04-29	female	hetero	Роза  Шмуклер Russia Moscow 	392	56.05396	37.170393	{}	confirmed	t	6
394	user393@gmail.com	fhjegga	Сергей	Колесник	\N	male	hetero	Сергей  Колесник Russia Moscow 	393	55.77896	37.345393	{culture,cooking,"drink beer","find something new"}	confirmed	t	9
395	user394@gmail.com	fhjegga	Антон	Бородин	\N	male	homo	Антон  Бородин Russia Moscow 	394	55.47896	37.595393	{"find something new"}	confirmed	t	9
396	user395@gmail.com	fhjegga	Марина	Мазураш	\N	female		Марина  Мазураш Russia Moscow 	395	55.85396	38.120393	{"video games",architecture}	confirmed	t	8
397	user396@gmail.com	fhjegga	Макс	Степанов	1978-08-01	male	hetero	Макс  Степанов Russia Moscow 	396	56.00396	37.620393	{architecture,programming,"find something new","video games"}	confirmed	t	7
398	user397@gmail.com	fhjegga	Татьяна	Васильева	\N	female	hetero	Татьяна  Васильева 	397	\N	\N	{football,politics}	confirmed	t	6
399	user398@gmail.com	fhjegga	Татьяна	Колос	1983-04-19	female	hetero	Татьяна  Колос Russia Moscow 	398	55.30396	37.970393	{}	confirmed	t	1
400	user399@gmail.com	fhjegga	Sasha	Gorbunov	\N	male	hetero	Sasha  Gorbunov Russia Moscow 	399	55.65396	37.970393	{}	confirmed	t	2
401	user400@gmail.com	fhjegga	Юрий	Спиридонов	1980-04-03	male	hetero	Юрий  Спиридонов Russia Moscow 	400	55.40396	37.795393	{"drink beer","video games",football}	confirmed	t	9
402	user401@gmail.com	fhjegga	Катерина	Афанасикова	\N		hetero	Катерина  Афанасикова Russia Moscow 	401	55.55396	37.470393	{}	confirmed	t	9
403	user402@gmail.com	fhjegga	Дмитрий	Ралко	1972-07-09	male	hetero	Дмитрий  Ралко Russia Moscow 	402	55.72896	37.695393	{programming}	confirmed	t	5
404	user403@gmail.com	fhjegga	Василий	Макагон	1979-12-07	male	hetero	Василий  Макагон Russia Moscow 	403	55.90396	37.520393	{football,"video games"}	confirmed	t	10
405	user404@gmail.com	fhjegga	Юлия	Прохорова	\N	female		Юлия  Прохорова Russia Moscow 	404	55.70396	38.020393	{"video games",programming,cooking}	confirmed	t	2
406	user405@gmail.com	fhjegga	Сергей	Корнилов	\N	male	hetero	Сергей  Корнилов Russia Moscow 	405	55.37896	37.645393	{}	confirmed	t	5
407	user406@gmail.com	fhjegga	Наталья	Артемьева	\N	female	hetero	Наталья  Артемьева 	406	\N	\N	{"drink beer",cooking}	confirmed	t	9
408	user407@gmail.com	fhjegga	Мел	Гибсон	1920-01-01		hetero	Мел  Гибсон 	407	\N	\N	{"find something new",architecture,cooking}	confirmed	t	1
409	user408@gmail.com	fhjegga	Александр	Чечко	1983-05-19	male	homo	Александр  Чечко Russia Moscow 	408	55.92896	37.545393	{}	confirmed	t	10
410	user409@gmail.com	fhjegga	Мария	Смирнова	\N	female	hetero	Мария  Борисова (Смирнова) Russia Moscow 	409	55.27896	37.120393	{politics,youtube,architecture}	confirmed	t	5
411	user410@gmail.com	fhjegga	Иван	Зарубенко	\N	male	hetero	Иван  Зарубенко Russia Moscow 	410	55.90396	37.670393	{politics}	confirmed	t	7
412	user411@gmail.com	fhjegga	Санек	Бермяков	1981-09-10		hetero	Санек  Бермяков Russia Moscow 	411	55.90396	37.920393	{architecture}	confirmed	t	2
413	user412@gmail.com	fhjegga	Ванек	Василюк	\N			Ванек  Василюк 	412	\N	\N	{"find something new"}	confirmed	t	10
414	user413@gmail.com	fhjegga	Дмитрий	Иванов	1987-09-26	male		Дмитрий  Иванов Russia Moscow 	413	55.75396	37.645393	{programming,"drink beer","find something new"}	confirmed	t	9
415	user414@gmail.com	fhjegga	Егор	Мешков	\N	male	hetero	Егор  Мешков Russia Moscow 	414	55.45396	37.595393	{programming,architecture,football,cooking}	confirmed	t	9
416	user415@gmail.com	fhjegga	Леня	Дрико	1988-12-12		homo	Леня  Дрико Russia Moscow 	415	55.92896	37.395393	{cooking}	confirmed	t	1
417	user416@gmail.com	fhjegga	Сергей	Маркевич	1990-08-27	male	hetero	Сергей  Маркевич Russia Moscow 	416	55.30396	37.220393	{football}	confirmed	t	4
418	user417@gmail.com	fhjegga	Ольга	Михальченко	1992-07-23	female	hetero	Ольга  Михальченко 	417	\N	\N	{programming,"drink beer"}	confirmed	t	2
419	user418@gmail.com	fhjegga	Ольга	Сигаева	\N	female	hetero	Ольга  Павловская (Сигаева) Russia Moscow 	418	55.85396	37.370393	{}	confirmed	t	10
420	user419@gmail.com	fhjegga	Костя	Хадиков	\N	male	homo	Костя  Хадиков 	419	\N	\N	{}	confirmed	t	6
421	user420@gmail.com	fhjegga	Сабина	Тайм	1987-06-20	female	homo	Сабина  Тайм Russia Moscow 	420	55.92896	37.520393	{"find something new"}	confirmed	t	5
422	user421@gmail.com	fhjegga	Илья	Гоов	\N	male	homo	Илья  Гоов Russia Moscow 	421	55.42896	37.770393	{youtube,football,"drink beer"}	confirmed	t	5
423	user422@gmail.com	fhjegga	Виктор	Баланин	\N	male	homo	Виктор  Баланин Russia Moscow 	422	56.25396	37.795393	{"drink beer","find something new",programming}	confirmed	t	7
424	user423@gmail.com	fhjegga	Егор	Дружинин	\N	male	hetero	Егор  Дружинин Russia Moscow 	423	55.65396	37.495393	{youtube,"find something new","video games"}	confirmed	t	6
425	user424@gmail.com	fhjegga	Дмитрий	Дерябин	\N	male	homo	Дмитрий  Дерябин Russia Moscow 	424	56.02896	37.295393	{}	confirmed	t	9
426	user425@gmail.com	fhjegga	Максим	Матюшенко	\N	male	hetero	Максим  Матюшенко 	425	\N	\N	{}	confirmed	t	9
427	user426@gmail.com	fhjegga	Шамиль	Шиллаев	1984-08-07	female	homo	Шамиль  Шиллаев Russia Moscow 	426	56.12896	37.770393	{"video games",politics}	confirmed	t	2
428	user427@gmail.com	fhjegga	Роман	Селлин	1984-10-01	male	homo	Роман  Селлин Russia Moscow 	427	56.22896	37.195393	{"drink beer","find something new","video games"}	confirmed	t	0
429	user428@gmail.com	fhjegga	Игорь	Печерица	\N	male	hetero	Игорь  Печерица Ukraine Kiev 	428	50.6036	30.0164	{}	confirmed	t	0
430	user429@gmail.com	fhjegga	Новая	Страница	\N	female	hetero	Новая  Страница Russia Moscow 	429	56.15396	37.445393	{architecture}	confirmed	t	8
431	user430@gmail.com	fhjegga	Влад	Парголовский	1983-11-18	female	homo	Влад  Парголовский 	430	\N	\N	{}	confirmed	t	0
432	user431@gmail.com	fhjegga	Lala	Mamedova	\N	female	hetero	Lala  Mamedova 	431	\N	\N	{youtube,architecture}	confirmed	t	3
433	user432@gmail.com	fhjegga	Данил	Zarembo	1990-12-08	male	hetero	Данил  Zarembo Russia Moscow 	432	55.50396	37.320393	{culture,"video games","find something new",architecture}	confirmed	t	3
434	user433@gmail.com	fhjegga	Julianna	Fogel	\N		hetero	Julianna  Fogel 	433	\N	\N	{"find something new","drink beer"}	confirmed	t	8
435	user434@gmail.com	fhjegga	Денис	Шихов	1988-01-15	male	hetero	Денис AR Шихов 	434	\N	\N	{}	confirmed	t	4
436	user435@gmail.com	fhjegga	Александр	Кузнецов	1988-05-14	male	hetero	Александр  Кузнецов Russia Moscow 	435	55.65396	37.545393	{football}	confirmed	t	4
437	user436@gmail.com	fhjegga	Михей	Шолохов	\N		homo	Михей  Шолохов 	436	\N	\N	{"drink beer","find something new",football}	confirmed	t	6
438	user437@gmail.com	fhjegga	Валерий	Михалев	1945-08-12	male	homo	Валерий  Михалев Russia Moscow 	437	56.15396	37.745393	{programming,cooking,football,architecture}	confirmed	t	0
439	user438@gmail.com	fhjegga	Валентин	Лысенко	\N	female	hetero	Валентин  Лысенко Russia Moscow 	438	56.17896	37.195393	{culture,cooking,football,"find something new"}	confirmed	t	1
440	user439@gmail.com	fhjegga	Tony	Montana	\N	female	hetero	Tony  Montana Russia Moscow 	439	55.82896	37.520393	{}	confirmed	t	6
441	user440@gmail.com	fhjegga	Олька	Дубинина	\N	female	hetero	Олька  Фунтусова (Дубинина) Russia Moscow 	440	55.32896	37.345393	{"drink beer",youtube,politics}	confirmed	t	4
442	user441@gmail.com	fhjegga	Наталья	Кузнецова	\N	female	homo	Наталья  Кузнецова Russia Moscow 	441	56.20396	37.545393	{architecture}	confirmed	t	7
443	user442@gmail.com	fhjegga	Сергей	Степанов	\N	male	hetero	Сергей  Степанов 	442	\N	\N	{culture,football}	confirmed	t	8
444	user443@gmail.com	fhjegga	Мария	Архангельская	\N	female	homo	Мария  Архангельская Russia Moscow 	443	55.62896	38.045393	{}	confirmed	t	0
445	user444@gmail.com	fhjegga	Екатерина	Куделич	\N	female	hetero	Екатерина  Куделич (Куделич) Russia Moscow 	444	55.37896	37.645393	{"video games",youtube,"drink beer",football}	confirmed	t	5
446	user445@gmail.com	fhjegga	Vovan	Darsavi	\N	male	homo	Vovan  Darsavi Russia Moscow 	445	56.07896	37.270393	{politics,programming}	confirmed	t	5
447	user446@gmail.com	fhjegga	Мария	Саракуца	\N	female	hetero	Мария  Саракуца Ukraine Kiev 	446	50.2036	30.741400000000002	{}	confirmed	t	7
448	user447@gmail.com	fhjegga	Nataly	Svetik	\N	male	homo	Nataly  Svetik Russia Moscow 	447	55.27896	38.070393	{"find something new"}	confirmed	t	5
449	user448@gmail.com	fhjegga	Игорь	Иванов	1986-03-05	male	hetero	Игорь  Иванов Russia Moscow 	448	55.67896	38.120393	{football,architecture}	confirmed	t	1
450	user449@gmail.com	fhjegga	Владимир	Ревенк	1985-07-10	male		Владимир  Ревенк Russia Moscow 	449	55.50396	37.945393	{cooking}	confirmed	t	11
451	user450@gmail.com	fhjegga	Юра	Оо	\N	male	homo	Юра  Оо Russia Moscow 	450	55.67896	37.795393	{architecture,"video games"}	confirmed	t	0
452	user451@gmail.com	fhjegga	Екатерина	Максимова	\N	female	hetero	Екатерина  Максимова Russia Moscow 	451	55.60396	37.145393	{culture,football}	confirmed	t	8
453	user452@gmail.com	fhjegga	Владимир	Четвертаков	1988-10-17	male	hetero	Владимир Юрьевич Четвертаков Russia Moscow 	452	56.02896	37.695393	{}	confirmed	t	3
454	user453@gmail.com	fhjegga	Павел	Журбиков	1980-06-20	male		Павел  Журбиков Russia Moscow 	453	55.32896	37.595393	{}	confirmed	t	5
455	user454@gmail.com	fhjegga	Осип	Высоцкий	1987-07-26			Осип  Высоцкий Russia Moscow 	454	56.07896	37.795393	{politics,cooking}	confirmed	t	8
456	user455@gmail.com	fhjegga	Александра	Филиппова	1990-03-12			Александра  Филиппова Russia Moscow 	455	55.67896	37.595393	{}	confirmed	t	8
457	user456@gmail.com	fhjegga	Елена	Шульгина	1987-06-25	female	hetero	Елена  Бриц (Шульгина) Russia Moscow 	456	55.32896	37.495393	{culture}	confirmed	t	6
458	user457@gmail.com	fhjegga	Анюточка	Шабанова	\N	female		Анюточка  Зорькина (Шабанова) Russia Moscow 	457	55.67896	37.670393	{architecture}	confirmed	t	3
459	user458@gmail.com	fhjegga	Ольга	Яглинская	\N	female	hetero	Ольга  Яглинская Russia Moscow 	458	56.12896	37.770393	{politics,programming}	confirmed	t	0
460	user459@gmail.com	fhjegga	Тимофей	Сапожников	\N	male	hetero	Тимофей Леонидович Сапожников Russia Moscow 	459	55.85396	37.495393	{youtube,programming,"drink beer","find something new"}	confirmed	t	1
461	user460@gmail.com	fhjegga	Noname	Noname	1989-11-09	female	hetero	Noname  Noname Russia Moscow 	460	55.82896	37.620393	{architecture,cooking,youtube}	confirmed	t	4
462	user461@gmail.com	fhjegga	Александр	Вакарчук	\N	male	homo	Александр  Вакарчук Russia Moscow 	461	55.62896	38.095393	{architecture,football}	confirmed	t	0
463	user462@gmail.com	fhjegga	Максим	Орленко	1982-02-07	male	hetero	Максим  Орленко Russia Moscow 	462	55.65396	37.845393	{culture}	confirmed	t	1
464	user463@gmail.com	fhjegga	Оксана	Король	1990-02-05	female	homo	Оксана  Король Russia Moscow 	463	55.95396	38.045393	{"video games",youtube}	confirmed	t	2
465	user464@gmail.com	fhjegga	Pavel	Steinberg	1988-12-20	female	hetero	Pavel  Steinberg 	464	\N	\N	{politics,architecture}	confirmed	t	0
466	user465@gmail.com	fhjegga	Димон	Симон	1983-06-29		hetero	Димон  Симон Russia Moscow 	465	56.07896	37.645393	{"drink beer",politics,"video games"}	confirmed	t	0
467	user466@gmail.com	fhjegga	Ruslan	Kalmykov	\N	male	hetero	Ruslan  Kalmykov Russia Moscow 	466	55.30396	37.420393	{architecture,youtube,"video games",politics}	confirmed	t	0
468	user467@gmail.com	fhjegga	Лиана	Магомедова	1989-01-31		homo	Лиана  Магомедова Russia Moscow 	467	56.12896	37.695393	{programming,"drink beer",youtube,cooking}	confirmed	t	9
469	user468@gmail.com	fhjegga	Светлина	Коршунова	\N	female	hetero	Светлина  Коршунова 	468	\N	\N	{football}	confirmed	t	6
470	user469@gmail.com	fhjegga	Константин	Алимов	1985-11-12	male	hetero	Константин  Алимов Russia Moscow 	469	55.50396	38.070393	{}	confirmed	t	11
471	user470@gmail.com	fhjegga	Иван	Грибов	1988-08-13	male	hetero	Иван  Грибов Russia Moscow 	470	56.10396	38.045393	{programming}	confirmed	t	10
472	user471@gmail.com	fhjegga	Дмитрий	Махин	\N	male	hetero	Дмитрий  Махин 	471	\N	\N	{}	confirmed	t	10
473	user472@gmail.com	fhjegga	Саша	Вохмянин	1987-08-05		hetero	Саша Vohmed Вохмянин Russia Moscow 	472	55.55396	37.870393	{"find something new",football}	confirmed	t	1
474	user473@gmail.com	fhjegga	Александр	Дмитриев	1987-07-28	male	hetero	Александр  Дмитриев Russia Moscow 	473	55.47896	37.170393	{cooking}	confirmed	t	11
475	user474@gmail.com	fhjegga	Сашка	Меньшенин	1983-02-04		homo	Сашка  Меньшенин Russia Moscow 	474	56.17896	37.120393	{"drink beer"}	confirmed	t	0
476	user475@gmail.com	fhjegga	Антон	Швыдченко	1988-06-12	male	hetero	Антон  Швыдченко Russia Moscow 	475	55.37896	37.670393	{"find something new",youtube}	confirmed	t	5
477	user476@gmail.com	fhjegga	Кристина	Белова	1985-03-26	female	homo	Кристина  Белова Russia Moscow 	476	55.87896	37.520393	{}	confirmed	t	6
478	user477@gmail.com	fhjegga	Михаил	Вишленков	1987-04-10	male	homo	Михаил  Вишленков Russia Moscow 	477	56.22896	37.670393	{cooking}	confirmed	t	5
479	user478@gmail.com	fhjegga	Денис	Соколов	1985-09-20	male	hetero	Денис  Соколов Russia Moscow 	478	55.35396	37.445393	{}	confirmed	t	2
480	user479@gmail.com	fhjegga	Нина	Власенко	\N	female	homo	Нина  Власенко Russia Moscow 	479	55.97896	37.220393	{}	confirmed	t	1
481	user480@gmail.com	fhjegga	Константин	Минин	1985-01-14	male	hetero	Константин  Минин Russia Moscow 	480	55.80396	37.795393	{cooking,"drink beer",architecture,youtube}	confirmed	t	1
482	user481@gmail.com	fhjegga	Deleted	Deleted	\N	male	homo	Deleted  Deleted 	481	\N	\N	{football,architecture,cooking}	confirmed	t	2
483	user482@gmail.com	fhjegga	Сережа	Овчаренко	\N	male	hetero	Сережа  Овчаренко 	482	\N	\N	{culture}	confirmed	t	6
484	user483@gmail.com	fhjegga	Ekaterina	Duginova	1990-10-10		hetero	Ekaterina  Duginova Russia Moscow 	483	56.17896	37.520393	{programming}	confirmed	t	2
485	user484@gmail.com	fhjegga	Ксения	Ларионова	1985-06-05	female	hetero	Ксения  Перова (Ларионова) Russia Moscow 	484	56.12896	38.045393	{youtube,programming}	confirmed	t	8
486	user485@gmail.com	fhjegga	Антон	Егоров	1987-06-12	male	hetero	Антон  Егоров Russia Moscow 	485	55.85396	37.420393	{cooking,politics}	confirmed	t	3
487	user486@gmail.com	fhjegga	Антон	Ефримов	1989-04-14	male	hetero	Антон  Ефримов Russia Moscow 	486	56.00396	37.670393	{}	confirmed	t	2
488	user487@gmail.com	fhjegga	Дмитрий	Романов	\N	male	homo	Дмитрий  Романов Russia Moscow 	487	55.40396	38.020393	{}	confirmed	t	4
489	user488@gmail.com	fhjegga	Ян	Бублик	1982-01-01	female	homo	Ян  Бублик Russia Moscow 	488	55.47896	37.845393	{cooking,politics}	confirmed	t	2
490	user489@gmail.com	fhjegga	Дмитрий	Ли	\N	male		Дмитрий  Ли Russia Moscow 	489	55.75396	38.095393	{cooking}	confirmed	t	2
491	user490@gmail.com	fhjegga	Константин	Прохоров	1988-02-07	male	hetero	Константин  Прохоров Russia Moscow 	490	56.17896	37.595393	{politics}	confirmed	t	3
492	user491@gmail.com	fhjegga	Андрей	Жигулин	\N	male	hetero	Андрей  Жигулин Russia Moscow 	491	56.25396	37.270393	{football}	confirmed	t	0
493	user492@gmail.com	fhjegga	Бу	Бу	\N		homo	Бу Bu Бу 	492	\N	\N	{"video games",architecture}	confirmed	t	1
494	user493@gmail.com	fhjegga	Оксана	Кузина	1986-11-29	female	homo	Оксана  Кузина Russia Moscow 	493	55.55396	37.895393	{programming,culture}	confirmed	t	9
495	user494@gmail.com	fhjegga	Наталья	Иванникова	\N	female	hetero	Наталья  Сидельникова (Иванникова) Russia Moscow 	494	55.95396	37.620393	{football,architecture,programming}	confirmed	t	6
496	user495@gmail.com	fhjegga	Сергей	Сеничев	\N	male	hetero	Сергей  Сеничев Russia Moscow 	495	55.47896	37.345393	{programming,cooking}	confirmed	t	6
497	user496@gmail.com	fhjegga	Женя	Шешуков	1983-08-04	male	hetero	Женя  Шешуков Russia Moscow 	496	55.82896	37.195393	{youtube,"drink beer",cooking,programming}	confirmed	t	11
498	user497@gmail.com	fhjegga	Сергей	Соловьев	\N	male	hetero	Сергей  Соловьев Russia Moscow 	497	55.55396	37.345393	{youtube,culture,programming,"video games"}	confirmed	t	8
499	user498@gmail.com	fhjegga	Патрик	Ефименко	1990-09-24	male	hetero	Патрик Протагонист Ефименко 	498	\N	\N	{architecture}	confirmed	t	1
500	user499@gmail.com	fhjegga	Alik	Зангиев	\N		hetero	Alik  Зангиев Russia Moscow 	499	55.37896	37.395393	{"find something new","video games",politics}	confirmed	t	1
501	user500@gmail.com	fhjegga	Ксаночка	Полякова	\N		hetero	Ксаночка  Полякова Russia Moscow 	500	55.42896	37.770393	{"video games"}	confirmed	t	5
\.


--
-- Name: devices_id_seq; Type: SEQUENCE SET; Schema: public; Owner: bsabre
--

SELECT pg_catalog.setval('public.devices_id_seq', 1, false);


--
-- Name: history_id_seq; Type: SEQUENCE SET; Schema: public; Owner: bsabre
--

SELECT pg_catalog.setval('public.history_id_seq', 1, false);


--
-- Name: interests_id_seq; Type: SEQUENCE SET; Schema: public; Owner: bsabre
--

SELECT pg_catalog.setval('public.interests_id_seq', 10, true);


--
-- Name: messages_mid_seq; Type: SEQUENCE SET; Schema: public; Owner: bsabre
--

SELECT pg_catalog.setval('public.messages_mid_seq', 1, false);


--
-- Name: notifs_nid_seq; Type: SEQUENCE SET; Schema: public; Owner: bsabre
--

SELECT pg_catalog.setval('public.notifs_nid_seq', 1, false);


--
-- Name: photos_pid_seq; Type: SEQUENCE SET; Schema: public; Owner: bsabre
--

SELECT pg_catalog.setval('public.photos_pid_seq', 500, true);


--
-- Name: users_uid_seq; Type: SEQUENCE SET; Schema: public; Owner: bsabre
--

SELECT pg_catalog.setval('public.users_uid_seq', 501, true);


--
-- Name: claims claims_pkey; Type: CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.claims
    ADD CONSTRAINT claims_pkey PRIMARY KEY (uidsender, uidreceiver);


--
-- Name: devices devices_pkey; Type: CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.devices
    ADD CONSTRAINT devices_pkey PRIMARY KEY (id);


--
-- Name: history history_pkey; Type: CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.history
    ADD CONSTRAINT history_pkey PRIMARY KEY (id);


--
-- Name: ignores ignores_pkey; Type: CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.ignores
    ADD CONSTRAINT ignores_pkey PRIMARY KEY (uidsender, uidreceiver);


--
-- Name: interests interests_pkey; Type: CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.interests
    ADD CONSTRAINT interests_pkey PRIMARY KEY (id);


--
-- Name: likes likes_pkey; Type: CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.likes
    ADD CONSTRAINT likes_pkey PRIMARY KEY (uidsender, uidreceiver);


--
-- Name: messages messages_pkey; Type: CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT messages_pkey PRIMARY KEY (mid);


--
-- Name: notifs notifs_pkey; Type: CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.notifs
    ADD CONSTRAINT notifs_pkey PRIMARY KEY (nid);


--
-- Name: photos photos_pkey; Type: CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.photos
    ADD CONSTRAINT photos_pkey PRIMARY KEY (pid);


--
-- Name: users users_mail_key; Type: CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_mail_key UNIQUE (mail);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (uid);


--
-- Name: birth_idx; Type: INDEX; Schema: public; Owner: bsabre
--

CREATE INDEX birth_idx ON public.users USING btree (birth);


--
-- Name: interests_idx; Type: INDEX; Schema: public; Owner: bsabre
--

CREATE INDEX interests_idx ON public.users USING btree (interests);


--
-- Name: interests_table_idx; Type: INDEX; Schema: public; Owner: bsabre
--

CREATE INDEX interests_table_idx ON public.interests USING btree (name);


--
-- Name: location_idx; Type: INDEX; Schema: public; Owner: bsabre
--

CREATE INDEX location_idx ON public.users USING btree (latitude, longitude);


--
-- Name: rating_idx; Type: INDEX; Schema: public; Owner: bsabre
--

CREATE INDEX rating_idx ON public.users USING btree (rating);


--
-- Name: sex_idx; Type: INDEX; Schema: public; Owner: bsabre
--

CREATE INDEX sex_idx ON public.users USING btree (gender, orientation);


--
-- Name: claims claims_receiver_fkey; Type: FK CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.claims
    ADD CONSTRAINT claims_receiver_fkey FOREIGN KEY (uidreceiver) REFERENCES public.users(uid) ON DELETE RESTRICT;


--
-- Name: claims claims_sender_fkey; Type: FK CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.claims
    ADD CONSTRAINT claims_sender_fkey FOREIGN KEY (uidsender) REFERENCES public.users(uid) ON DELETE RESTRICT;


--
-- Name: devices device_fkey; Type: FK CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.devices
    ADD CONSTRAINT device_fkey FOREIGN KEY (uid) REFERENCES public.users(uid) ON DELETE RESTRICT;


--
-- Name: history history_target_uid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.history
    ADD CONSTRAINT history_target_uid_fkey FOREIGN KEY (targetuid) REFERENCES public.users(uid) ON DELETE RESTRICT;


--
-- Name: history history_uid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.history
    ADD CONSTRAINT history_uid_fkey FOREIGN KEY (uid) REFERENCES public.users(uid) ON DELETE RESTRICT;


--
-- Name: ignores ignores_receiver_fkey; Type: FK CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.ignores
    ADD CONSTRAINT ignores_receiver_fkey FOREIGN KEY (uidreceiver) REFERENCES public.users(uid) ON DELETE RESTRICT;


--
-- Name: ignores ignores_sender_fkey; Type: FK CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.ignores
    ADD CONSTRAINT ignores_sender_fkey FOREIGN KEY (uidsender) REFERENCES public.users(uid) ON DELETE RESTRICT;


--
-- Name: likes likereceiver_fkey; Type: FK CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.likes
    ADD CONSTRAINT likereceiver_fkey FOREIGN KEY (uidreceiver) REFERENCES public.users(uid) ON DELETE RESTRICT;


--
-- Name: likes likesender_fkey; Type: FK CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.likes
    ADD CONSTRAINT likesender_fkey FOREIGN KEY (uidsender) REFERENCES public.users(uid) ON DELETE RESTRICT;


--
-- Name: notifs notifreceiver_fkey; Type: FK CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.notifs
    ADD CONSTRAINT notifreceiver_fkey FOREIGN KEY (uidreceiver) REFERENCES public.users(uid) ON DELETE RESTRICT;


--
-- Name: notifs notifsender_fkey; Type: FK CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.notifs
    ADD CONSTRAINT notifsender_fkey FOREIGN KEY (uidsender) REFERENCES public.users(uid) ON DELETE RESTRICT;


--
-- Name: photos photos_fkey; Type: FK CONSTRAINT; Schema: public; Owner: bsabre
--

ALTER TABLE ONLY public.photos
    ADD CONSTRAINT photos_fkey FOREIGN KEY (uid) REFERENCES public.users(uid) ON DELETE RESTRICT;


--
-- PostgreSQL database dump complete
--

