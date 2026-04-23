package main

// indexTmpl - template for index page
const indexTmpl = `<!DOCTYPE html>
<!-- Served by shell2http/%s -->
<html data-theme="dark"%s>
<head>
    <title>❯ CrowdStrike's VulnApp</title>
    <link rel="icon" type="image/png" sizes="32x32" href="/images/favicon-32.png">
    <link rel="icon" type="image/png" sizes="16x16" href="/images/favicon-16.png">
    <link rel="apple-touch-icon" title="CrowdStrike" href="/images/logo.png">
    <style>
    /* ── CSS Custom Properties (Dark / Light) ── */
    :root, [data-theme="dark"] {
        --bg-primary: #0a0a0f;
        --bg-secondary: #12121a;
        --bg-card: #1a1a25;
        --bg-card-hover: #222233;
        --text-primary: rgba(255,255,255,0.87);
        --text-secondary: rgba(255,255,255,0.65);
        --accent: #e8263a;
        --accent-glow: rgba(232,38,58,0.4);
        --accent-cyan: #00f0ff;
        --border: rgba(255,255,255,0.08);
        --scanline: rgba(0,240,255,0.03);
        --grid-color: rgba(0,240,255,0.07);
    }
    [data-theme="light"] {
        --bg-primary: #f0f0f5;
        --bg-secondary: #e8e8ef;
        --bg-card: #ffffff;
        --bg-card-hover: #f5f5fa;
        --text-primary: rgba(0,0,0,0.87);
        --text-secondary: rgba(0,0,0,0.6);
        --accent: #c8102e;
        --accent-glow: rgba(200,16,46,0.25);
        --accent-cyan: #007a8a;
        --border: rgba(0,0,0,0.1);
        --scanline: rgba(0,0,0,0.02);
        --grid-color: rgba(0,122,138,0.06);
    }

    /* ── Reset & Base ── */
    *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }
    html { overflow-x: hidden; }
    body {
        font-family: 'Courier New', 'Consolas', monospace;
        background-color: var(--bg-primary);
        color: var(--text-primary);
        min-height: 100vh;
        overflow-x: hidden;
        position: relative;
    }

    /* ── Animated Grid Background ── */
    body::before {
        content: '';
        position: fixed;
        inset: 0;
        background-image:
            linear-gradient(var(--grid-color) 1px, transparent 1px),
            linear-gradient(90deg, var(--grid-color) 1px, transparent 1px);
        background-size: 60px 60px;
        animation: gridScroll 20s linear infinite;
        z-index: 0;
        pointer-events: none;
    }

    /* ── Scanline Overlay ── */
    body::after {
        content: '';
        position: fixed;
        inset: 0;
        background: repeating-linear-gradient(
            0deg,
            var(--scanline),
            var(--scanline) 1px,
            transparent 1px,
            transparent 3px
        );
        z-index: 9998;
        pointer-events: none;
    }

    /* ── Sweeping CRT Scanline ── */
    #sweep-scanline {
        position: fixed;
        left: 0;
        width: 100%%;
        height: 6px;
        z-index: 9997;
        pointer-events: none;
        background: linear-gradient(
            180deg,
            transparent,
            rgba(0,240,255,0.08) 20%%,
            rgba(0,240,255,0.15) 50%%,
            rgba(0,240,255,0.08) 80%%,
            transparent
        );
        box-shadow: 0 0 20px 4px rgba(0,240,255,0.06);
        animation: none;
    }
    [data-theme="light"] #sweep-scanline {
        background: linear-gradient(
            180deg,
            transparent,
            rgba(0,122,138,0.06) 20%%,
            rgba(0,122,138,0.12) 50%%,
            rgba(0,122,138,0.06) 80%%,
            transparent
        );
        box-shadow: 0 0 20px 4px rgba(0,122,138,0.04);
    }
    @keyframes sweepDown {
        0%% { top: -6px; opacity: 1; }
        85%% { opacity: 1; }
        100%% { top: 100vh; opacity: 0; }
    }

    @keyframes gridScroll {
        0%% { transform: translate(0, 0); }
        100%% { transform: translate(60px, 60px); }
    }

    /* ── Header ── */
    .header-bar {
        position: relative;
        z-index: 10;
        display: flex;
        align-items: center;
        padding: 12px 24px;
        background: var(--bg-secondary);
        border-bottom: 1px solid var(--border);
        box-shadow: 0 2px 20px var(--accent-glow);
    }
    .header-bar img.logo {
        height: 32px;
        animation: logoPulse 3s ease-in-out infinite;
    }
    .header-bar .separator {
        color: var(--text-secondary);
        padding: 0 12px;
        font-size: 1.2rem;
        align-self: flex-start;
        line-height: 32px;
    }
    .header-bar h2 {
        font-size: 1.2rem;
        letter-spacing: 3px;
        text-transform: uppercase;
        color: var(--accent);
        text-shadow: 0 0 10px var(--accent-glow);
        align-self: flex-start;
        line-height: 32px;
    }

    @keyframes logoPulse {
        0%%, 100%% { filter: drop-shadow(0 0 4px var(--accent-glow)); }
        50%% { filter: drop-shadow(0 0 16px var(--accent)); }
    }

    /* ── Theme Toggle ── */
    .theme-toggle {
        margin-left: auto;
        background: transparent;
        border: 1px solid var(--border);
        color: var(--text-primary);
        font-size: 1.3rem;
        padding: 6px 10px;
        border-radius: 6px;
        cursor: pointer;
        transition: all 0.3s ease;
        z-index: 10;
    }
    .theme-toggle:hover {
        border-color: var(--accent);
        box-shadow: 0 0 12px var(--accent-glow);
        transform: scale(1.1);
    }

    /* ── Content ── */
    .content { position: relative; z-index: 1; padding: 0 24px 40px; }

    .container {
        display: grid;
        grid-template-columns: 1fr auto 1fr;
        gap: 24px;
        padding: 32px 0;
        align-items: center;
    }
    .welcome {
        grid-column: 1;
        max-width: 620px;
        justify-self: end;
    }
    .hero {
        width: 100%%;
        max-height: 460px;
        object-fit: contain;
        border-radius: 8px;
        opacity: 0.85;
        transition: opacity 0.5s;
    }
    .hero:hover { opacity: 1; }
    .hero-wrap {
        position: relative;
        grid-column: 2;
        max-width: 620px;
        overflow: hidden;
        border-radius: 8px;
    }
    .hero-wrap .grain-overlay {
        position: absolute;
        inset: 0;
        border-radius: 8px;
        z-index: 4;
        pointer-events: none;
        opacity: 0;
        mix-blend-mode: overlay;
        background-image: url("data:image/svg+xml,%%3Csvg xmlns='http://www.w3.org/2000/svg' width='200' height='200'%%3E%%3Cfilter id='n'%%3E%%3CfeTurbulence type='fractalNoise' baseFrequency='0.9' numOctaves='4' stitchTiles='stitch'/%%3E%%3C/filter%%3E%%3Crect width='100%%25' height='100%%25' filter='url(%%23n)' opacity='1'/%%3E%%3C/svg%%3E");
        background-size: 200px 200px;
    }
    .hero-wrap .grain-overlay.active {
        opacity: 0.45;
        animation: grainShift 0.15s steps(3) infinite;
    }
    @keyframes grainShift {
        0%%   { transform: translate(0, 0); }
        33%%  { transform: translate(-2px, 3px); }
        66%%  { transform: translate(3px, -2px); }
        100%% { transform: translate(-1px, 1px); }
    }
    .hero-ghost {
        position: absolute;
        inset: 0;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 12vw;
        color: var(--accent);
        text-shadow: 0 0 40px var(--accent-glow), 0 0 80px rgba(232,38,58,0.3);
        filter: blur(2px);
        opacity: 0;
        pointer-events: none;
        transition: opacity 0.15s;
    }
    .hero-adversary {
        position: absolute;
        inset: 0;
        width: 100%%;
        height: 100%%;
        object-fit: contain;
        border-radius: 8px;
        opacity: 0;
        pointer-events: none;
        z-index: 2;
    }
    .hero-adversary.slam-in {
        animation: adversarySlamIn 0.25s cubic-bezier(0.22, 1, 0.36, 1) forwards;
    }
    .hero-adversary.holding {
        opacity: 1;
        box-shadow: 0 0 30px var(--accent-glow), 0 0 60px rgba(232,38,58,0.25), inset 0 0 30px rgba(232,38,58,0.1);
        animation: adversaryMenace 1.2s ease-in-out infinite;
    }
    @keyframes adversaryMenace {
        0%%   { filter: contrast(1.1) saturate(1.2) brightness(0.95); transform: skewX(-0.5deg); }
        25%%  { filter: contrast(1.3) saturate(1.4) hue-rotate(-5deg) brightness(0.9); transform: skewX(0.3deg) translateX(1px); }
        50%%  { filter: contrast(1.1) saturate(1.1) brightness(1.05); transform: skewX(-0.3deg); }
        75%%  { filter: contrast(1.4) saturate(1.5) hue-rotate(5deg) brightness(0.85); transform: skewX(0.5deg) translateX(-1px); }
        100%% { filter: contrast(1.1) saturate(1.2) brightness(0.95); transform: skewX(-0.5deg); }
    }
    .hero-adversary.glitch-out {
        animation: adversaryGlitchOut 0.4s ease-in forwards;
    }

    /* Hero corruption phase — applied before adversary appears */
    .hero.corrupting {
        animation: heroCorrupt 0.5s steps(4);
    }
    @keyframes heroCorrupt {
        0%%   { filter: none; clip-path: inset(0); transform: none; }
        15%%  { filter: saturate(3) hue-rotate(20deg); clip-path: inset(10%% 0 30%% 0); transform: skewX(3deg); }
        30%%  { filter: saturate(0.2) brightness(1.6); clip-path: inset(50%% 0 10%% 0); transform: skewX(-5deg) translateX(4px); }
        50%%  { filter: hue-rotate(-40deg) contrast(2) brightness(0.6); clip-path: inset(5%% 0 60%% 0); transform: skewX(6deg) translateX(-6px); }
        70%%  { filter: invert(0.8) saturate(4); clip-path: inset(40%% 0 20%% 0); transform: skewX(-4deg) translateX(3px); }
        85%%  { filter: brightness(2) hue-rotate(60deg); clip-path: inset(15%% 0 45%% 0); transform: skewX(2deg); }
        100%% { filter: brightness(3) saturate(0); clip-path: inset(0); transform: scale(0.98); opacity: 0.3; }
    }

    /* Bright flash between corruption and reveal */
    .hero-flash {
        position: absolute;
        inset: 0;
        border-radius: 8px;
        opacity: 0;
        pointer-events: none;
        z-index: 3;
        background: radial-gradient(ellipse at center, rgba(232,38,58,0.9) 0%%, rgba(255,255,255,0.95) 40%%, transparent 70%%);
    }
    .hero-flash.active {
        animation: flashBurst 0.2s ease-out forwards;
    }
    @keyframes flashBurst {
        0%%   { opacity: 0; transform: scale(0.8); }
        40%%  { opacity: 1; transform: scale(1.1); }
        100%% { opacity: 0; transform: scale(1.3); }
    }

    /* Adversary slams in — fast scale from center with red glow */
    @keyframes adversarySlamIn {
        0%%   { opacity: 0; transform: scale(1.3); filter: brightness(2) blur(2px); }
        60%%  { opacity: 1; transform: scale(0.97); filter: brightness(1.1) blur(0); }
        100%% { opacity: 1; transform: none; filter: none; }
    }

    /* Container shake while adversary is entering */
    .hero-wrap.shaking {
        animation: containerShake 0.3s ease-out;
    }
    @keyframes containerShake {
        0%%   { transform: translate(0); }
        15%%  { transform: translate(-3px, 2px); }
        30%%  { transform: translate(4px, -2px); }
        45%%  { transform: translate(-2px, -1px); }
        60%%  { transform: translate(3px, 1px); }
        75%%  { transform: translate(-1px, 2px); }
        100%% { transform: translate(0); }
    }

    @keyframes adversaryGlitchOut {
        0%%   { opacity: 1; transform: none; filter: none; box-shadow: 0 0 30px var(--accent-glow); clip-path: inset(0); }
        10%%  { opacity: 1; transform: skewX(4deg) translateX(-6px); filter: hue-rotate(30deg) saturate(2) brightness(1.4); clip-path: inset(5%% 0 15%% 0); }
        25%%  { opacity: 0.9; transform: skewX(-6deg) translateX(8px); filter: hue-rotate(-50deg) brightness(1.6); clip-path: inset(30%% 0 10%% 0); }
        40%%  { opacity: 0.7; transform: skewX(8deg) translateX(-4px) scaleY(0.9); filter: invert(0.3) brightness(1.8) blur(2px); clip-path: inset(10%% 0 40%% 0); }
        55%%  { opacity: 0.5; transform: skewX(-10deg) scaleY(0.75); filter: hue-rotate(70deg) brightness(2) blur(3px); clip-path: inset(50%% 0 5%% 0); }
        70%%  { opacity: 0.3; transform: skewX(6deg) scaleY(0.5) translateX(6px); filter: saturate(4) brightness(2.5) blur(5px); clip-path: inset(20%% 0 50%% 0); }
        85%%  { opacity: 0.15; transform: skewX(-12deg) scaleY(0.3); filter: brightness(3) blur(6px); clip-path: inset(60%% 0 20%% 0); }
        100%% { opacity: 0; transform: skewX(15deg) scaleY(0.1); filter: brightness(4) blur(10px); clip-path: inset(0); }
    }
    .hero-ghost.visible {
        animation: heroGhostFlash 1.8s ease-out forwards;
    }
    @keyframes heroGhostFlash {
        0%% { opacity: 0; transform: scale(0.7); filter: blur(8px); }
        10%% { opacity: 0.35; transform: scale(1.1) skewX(-3deg); filter: blur(2px); }
        20%% { opacity: 0.1; transform: scale(0.9) skewX(3deg); filter: blur(5px); }
        35%% { opacity: 0.3; transform: scale(1.05); filter: blur(2px); }
        50%% { opacity: 0.12; transform: scale(0.97); filter: blur(4px); }
        65%% { opacity: 0.2; transform: scale(1.02); filter: blur(3px); }
        80%% { opacity: 0.06; filter: blur(6px); }
        100%% { opacity: 0; filter: blur(10px); }
    }

    /* ── Typing Effect on H1 ── */
    .welcome h1 {
        font-size: 1.6rem;
        color: var(--accent);
        border-right: 2px solid var(--accent);
        white-space: nowrap;
        overflow: hidden;
        width: fit-content;
        max-width: 0;
        animation: typeWriter 2.5s steps(36) 0.5s forwards, blinkCursor 0.7s step-end infinite;
    }
    @keyframes typeWriter {
        to { max-width: 100%%; }
    }
    @keyframes blinkCursor {
        50%% { border-color: transparent; }
    }

    .welcome p {
        color: var(--text-secondary);
        line-height: 1.6;
        margin-top: 12px;
        animation: fadeInUp 0.6s ease both;
    }
    .welcome p:nth-child(2) { animation-delay: 0.3s; }
    .welcome p:nth-child(3) { animation-delay: 0.5s; }
    .welcome p:nth-child(4) { animation-delay: 0.7s; }

    @keyframes fadeInUp {
        from { opacity: 0; transform: translateY(12px); }
        to { opacity: 1; transform: translateY(0); }
    }

    .welcome a {
        color: var(--accent-cyan);
        text-decoration: none;
        border-bottom: 1px dashed var(--accent-cyan);
        transition: all 0.2s;
    }
    .welcome a:hover {
        color: var(--accent);
        border-bottom-color: var(--accent);
        text-shadow: 0 0 6px var(--accent-glow);
    }

    /* ── Section Header ── */
    .section-header {
        font-size: 1.1rem;
        letter-spacing: 4px;
        text-transform: uppercase;
        color: var(--accent);
        margin: 32px 0 16px;
        padding-bottom: 8px;
        border-bottom: 1px solid var(--border);
        display: flex;
        align-items: center;
        gap: 8px;
    }
    .section-header::before {
        content: '>';
        color: var(--accent-cyan);
        animation: blinkCursor 1s step-end infinite;
    }

    /* ── Detection Cards ── */
    ul { list-style: none; display: grid; grid-template-columns: repeat(auto-fill, minmax(340px, 1fr)); gap: 12px; }
    li {
        background: var(--bg-card);
        border: 1px solid var(--border);
        border-left: 3px solid var(--accent);
        border-radius: 4px;
        padding: 8px 16px;
        transition: all 0.3s ease;
        position: relative;
        overflow: hidden;
        cursor: pointer;
    }
    li::before { content: none; }
    li::after {
        content: '';
        position: absolute;
        top: 0; left: -100%%;
        width: 100%%; height: 100%%;
        background: linear-gradient(90deg, transparent, var(--accent-glow), transparent);
        transition: left 0.5s ease;
    }
    li:hover {
        background: var(--bg-card-hover);
        border-left-color: var(--accent-cyan);
        transform: translateX(4px);
        box-shadow: 0 4px 20px rgba(0,0,0,0.3), -3px 0 12px var(--accent-glow);
    }
    li:hover::after { left: 100%%; }

    li a {
        color: var(--text-primary);
        text-decoration: none;
        font-weight: bold;
        font-size: 0.95rem;
    }
    li a::after {
        content: '';
        position: absolute;
        inset: 0;
    }

    /* ── Glitch Effect ── */
    li.random-glitch {
        animation: glitchHard 0.7s ease-in-out;
        border-left-color: var(--accent-cyan);
        box-shadow: 0 0 20px var(--accent-glow), 0 0 40px rgba(0,240,255,0.15);
    }
    li.random-glitch:hover {
        animation: none;
    }
    @keyframes glitchHard {
        0%% { transform: translate(0); }
        10%% { transform: translate(-4px, 2px) skewX(2deg); filter: hue-rotate(60deg) brightness(1.3); }
        20%% { transform: translate(3px, -3px); clip-path: inset(20%% 0 40%% 0); }
        30%% { transform: translate(-3px, 1px) skewX(-2deg); filter: hue-rotate(-50deg) saturate(1.8); }
        40%% { transform: translate(4px, -1px); clip-path: inset(60%% 0 5%% 0); }
        50%% { transform: translate(-2px, -2px); filter: hue-rotate(80deg) brightness(1.4); }
        60%% { transform: translate(3px, 2px) skewX(-2deg); clip-path: inset(10%% 0 50%% 0); }
        75%% { transform: translate(-3px, 1px); filter: hue-rotate(-60deg); }
        88%% { transform: translate(1px, -1px); }
        100%% { transform: translate(0); clip-path: inset(0); filter: none; }
    }
    @keyframes glitchText {
        0%% { text-shadow: none; }
        15%% { text-shadow: 2px 0 var(--accent-cyan), -2px 0 var(--accent); }
        30%% { text-shadow: -3px 0 var(--accent-cyan), 3px 0 var(--accent); letter-spacing: 2px; }
        50%% { text-shadow: 1px 1px var(--accent), -1px -1px var(--accent-cyan); }
        70%% { text-shadow: -2px 0 var(--accent-cyan), 2px 0 var(--accent); letter-spacing: 0; }
        100%% { text-shadow: none; }
    }
    li.random-glitch a { animation: glitchText 0.7s ease-in-out; }

    li span { color: rgba(255,255,255,0.75); font-size: 0.82rem; display: block; margin-top: 6px; }
    [data-theme="light"] li span { color: rgba(0,0,0,0.65); }

    /* ── Modal Overlay ── */
    #modal-overlay {
        position: fixed;
        inset: 0;
        z-index: 10000;
        background: rgba(0,0,0,0.75);
        backdrop-filter: blur(4px);
        display: none;
        align-items: center;
        justify-content: center;
        padding: 24px;
    }
    #modal-overlay.open { display: flex; }
    .modal-box {
        position: relative;
        background: var(--bg-secondary);
        border: 1px solid var(--accent);
        border-radius: 8px;
        box-shadow: 0 0 40px var(--accent-glow), 0 0 80px rgba(0,0,0,0.5);
        width: 100%%;
        max-width: 700px;
        max-height: 80vh;
        display: flex;
        flex-direction: column;
        overflow: hidden;
        animation: modalIn 0.3s ease-out;
    }
    @keyframes modalIn {
        from { opacity: 0; transform: scale(0.92) translateY(10px); }
        to { opacity: 1; transform: scale(1) translateY(0); }
    }
    .modal-header {
        display: flex;
        align-items: center;
        justify-content: space-between;
        padding: 14px 20px;
        border-bottom: 1px solid var(--border);
        background: var(--bg-card);
    }
    .modal-header h3 {
        font-size: 0.95rem;
        color: var(--accent-cyan);
        letter-spacing: 1px;
        margin: 0;
    }
    .modal-close {
        background: transparent;
        border: 1px solid var(--border);
        color: var(--text-primary);
        font-size: 1.2rem;
        width: 32px;
        height: 32px;
        border-radius: 4px;
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
        transition: all 0.2s;
    }
    .modal-close:hover {
        border-color: var(--accent);
        background: var(--accent);
        color: #fff;
        box-shadow: 0 0 12px var(--accent-glow);
    }
    .modal-body {
        padding: 20px;
        overflow-y: auto;
        font-size: 0.85rem;
        line-height: 1.7;
        color: rgba(255,255,255,0.75);
        white-space: pre-wrap;
        font-family: 'Courier New', monospace;
    }
    [data-theme="light"] .modal-body { color: rgba(0,0,0,0.65); }
    .modal-body .loading {
        color: var(--accent-cyan);
        animation: blinkCursor 0.7s step-end infinite;
    }

    /* ── Animations Disabled ── */
    [data-animations="off"] *,
    [data-animations="off"] *::before,
    [data-animations="off"] *::after {
        animation: none !important;
        transition: none !important;
    }
    [data-animations="off"] .welcome h1 {
        max-width: 100%% !important;
    }
    [data-animations="off"] .welcome p {
        opacity: 1 !important;
        transform: none !important;
    }
    [data-animations="off"] #matrix-canvas,
    [data-animations="off"] #sweep-scanline,
    [data-animations="off"] body::before,
    [data-animations="off"] body::after {
        display: none !important;
    }

    /* ── Animations Toggle ── */
    .anim-toggle {
        display: flex;
        align-items: center;
        gap: 6px;
        margin-left: 16px;
        cursor: pointer;
        font-size: 0.75rem;
        color: var(--text-secondary);
        user-select: none;
    }
    .anim-toggle input {
        accent-color: var(--accent);
        cursor: pointer;
    }

    /* ── Matrix Rain Canvas ── */
    #matrix-canvas {
        position: fixed;
        top: 0; left: 0;
        width: 100%%; height: 100%%;
        z-index: 0;
        pointer-events: none;
        opacity: 0.12;
    }
    [data-theme="light"] #matrix-canvas { opacity: 0.06; }
    </style>
</head>
<body>
    <canvas id="matrix-canvas"></canvas>
    <div id="sweep-scanline"></div>

    <div id="modal-overlay">
        <div class="modal-box">
            <div class="modal-header">
                <h3 id="modal-title"></h3>
                <button class="modal-close" onclick="closeModal()" aria-label="Close">✕</button>
            </div>
            <div class="modal-body" id="modal-body"></div>
        </div>
    </div>

    <div id="page-wrap">
    <div class="header-bar">
        <img class="logo" src="images/logo_crowdstrike.png" alt="CrowdStrike">
        <span class="separator">|</span>
        <h2>VulnApp</h2>
        <button class="theme-toggle" onclick="toggleTheme()" aria-label="Toggle theme">☾</button>
        <label class="anim-toggle"><input type="checkbox" id="anim-checkbox" onchange="toggleAnimations()"> Disable Animations</label>
    </div>

    <div class="content">
        <div class="container">
            <div class="welcome">
                <h1>Welcome to vulnerable.example.com</h1>
                <p>This web application runs on bare metal, virtual machines, or Kubernetes clusters utilizing CrowdStrike's Falcon sensor. All that is needed is a container runtime such as Docker or Podman. For Kubernetes, the Falcon sensor should be deployed as a DaemonSet or sidecar container.</p>
                <p>The web application will allow you to execute various exploitation techniques as if it was an attacker exploiting the application. The Falcon sensor will recognize this malicious behavior and report it back to the Falcon Console.</p>
                <p>You can view output of <a href="/ps">ps command</a> to see processes running within the same pod as this application.</p>
            </div>
            <div class="hero-wrap" id="hero-wrap">
                <img class="hero" src="images/hero-homepage.png" alt="VulnApp Hero">
                <img class="hero-adversary" id="hero-adversary" src="" alt="">
                <div class="grain-overlay" id="hero-grain"></div>
                <div class="hero-flash" id="hero-flash"></div>
                <div class="hero-ghost" id="hero-ghost">&#9760;</div>
            </div>
        </div>

        <div class="section-header">Detections</div>
        <ul>
            %s
        </ul>
    </div>
    </div>

    <script>
    // Theme toggle with localStorage persistence
    (function() {
        var saved = localStorage.getItem('vulnapp-theme');
        if (saved) {
            document.documentElement.setAttribute('data-theme', saved);
        }
        updateToggleIcon();
    })();

    function toggleTheme() {
        var current = document.documentElement.getAttribute('data-theme');
        var next = current === 'dark' ? 'light' : 'dark';
        document.documentElement.setAttribute('data-theme', next);
        localStorage.setItem('vulnapp-theme', next);
        updateToggleIcon();
    }

    function updateToggleIcon() {
        var btn = document.querySelector('.theme-toggle');
        if (btn) {
            btn.textContent = document.documentElement.getAttribute('data-theme') === 'dark' ? '☾' : '☀';
        }
    }

    // Animations toggle with localStorage persistence
    (function() {
        var saved = localStorage.getItem('vulnapp-animations');
        var cb = document.getElementById('anim-checkbox');
        if (saved === 'off') {
            document.documentElement.setAttribute('data-animations', 'off');
            if (cb) cb.checked = true;
        }
    })();

    function toggleAnimations() {
        var cb = document.getElementById('anim-checkbox');
        var disabled = cb && cb.checked;
        if (disabled) {
            document.documentElement.setAttribute('data-animations', 'off');
            localStorage.setItem('vulnapp-animations', 'off');
        } else {
            document.documentElement.removeAttribute('data-animations');
            localStorage.setItem('vulnapp-animations', 'on');
        }
    }

    // Matrix rain effect
    (function() {
        var canvas = document.getElementById('matrix-canvas');
        if (!canvas) return;
        var ctx = canvas.getContext('2d');
        canvas.width = window.innerWidth;
        canvas.height = window.innerHeight;

        var cols = Math.floor(canvas.width / 18);
        var drops = [];
        for (var i = 0; i < cols; i++) drops[i] = Math.random() * -100;

        var chars = '01アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲン';

        function draw() {
            ctx.fillStyle = 'rgba(10, 10, 15, 0.08)';
            ctx.fillRect(0, 0, canvas.width, canvas.height);
            ctx.fillStyle = '#e8263a';
            ctx.font = '14px monospace';

            for (var i = 0; i < drops.length; i++) {
                var ch = chars[Math.floor(Math.random() * chars.length)];
                ctx.fillText(ch, i * 18, drops[i] * 18);
                if (drops[i] * 18 > canvas.height && Math.random() > 0.975) {
                    drops[i] = 0;
                }
                drops[i]++;
            }
        }

        setInterval(draw, 50);
        window.addEventListener('resize', function() {
            canvas.width = window.innerWidth;
            canvas.height = window.innerHeight;
        });
    })();

    // Random glitch on detection cards
    (function() {
        function triggerRandomGlitch() {
            if (document.documentElement.getAttribute('data-animations') === 'off') return;
            var cards = document.querySelectorAll('li');
            if (cards.length === 0) return;
            var count = Math.random() < 0.3 ? 2 : 1;
            for (var n = 0; n < count; n++) {
                var card = cards[Math.floor(Math.random() * cards.length)];
                card.classList.add('random-glitch');
                (function(c) {
                    setTimeout(function() { c.classList.remove('random-glitch'); }, 700);
                })(card);
            }
        }
        setInterval(function() {
            if (Math.random() < 0.5) triggerRandomGlitch();
        }, 2200);
    })();

    // Hero image glitch effect
    (function() {
        var hero = document.querySelector('.hero');
        var ghost = document.getElementById('hero-ghost');
        var adversary = document.getElementById('hero-adversary');
        var symbols = ['☠', '⚠', '☣', '⛔', '☢'];
        var adversaryPaths = [
            'images/Bounty-Jackal-card.png',
            'images/Chatty-Spider.png',
            'images/Punk-Spider-card.png',
            'images/liminal-panda-card.png'
        ];
        var preloaded = [];
        adversaryPaths.forEach(function(src) {
            var img = new Image();
            img.src = src;
            preloaded.push(img);
        });
        var heroWrap = document.getElementById('hero-wrap');
        var flash = document.getElementById('hero-flash');
        var grain = document.getElementById('hero-grain');
        var adversaryActive = false;
        var adversaryIndex = 0;
        function heroGlitch() {
            if (document.documentElement.getAttribute('data-animations') === 'off') return;
            if (!hero || adversaryActive) return;
            var showAdversary = Math.random() < 0.5;
            var skew = (Math.random() * 6 - 3);
            hero.style.transition = 'none';
            hero.style.transform = 'skewX(' + skew + 'deg) scale(1.02)';
            hero.style.filter = 'hue-rotate(40deg) saturate(1.5) brightness(1.3)';
            if (showAdversary && adversary) {
                adversaryActive = true;
                var idx = Math.floor(Math.random() * preloaded.length);
                if (preloaded.length > 1 && idx === adversaryIndex) {
                    idx = (idx + 1) %% preloaded.length;
                }
                adversaryIndex = idx;
                var pick = preloaded[idx];
                adversary.src = pick.src;
                if (ghost) ghost.classList.remove('visible');

                // Phase 1: Corrupt the hero (500ms)
                hero.classList.add('corrupting');

                // Phase 2: Flash burst + slam adversary in (at 500ms)
                setTimeout(function() {
                    hero.classList.remove('corrupting');
                    hero.style.transition = 'none';
                    hero.style.opacity = '0';
                    hero.style.transform = '';
                    hero.style.filter = '';
                    if (flash) {
                        flash.classList.add('active');
                        setTimeout(function() { flash.classList.remove('active'); }, 200);
                    }
                    if (heroWrap) {
                        heroWrap.classList.add('shaking');
                        setTimeout(function() { heroWrap.classList.remove('shaking'); }, 300);
                    }
                    adversary.className = 'hero-adversary slam-in';
                    if (grain) grain.classList.add('active');

                    // Phase 3: Hold with menace (at 750ms, hold for 2s)
                    setTimeout(function() {
                        adversary.className = 'hero-adversary holding';
                    }, 250);
                }, 500);

                // Phase 4: Glitch out + restore hero (at 2750ms)
                setTimeout(function() {
                    adversary.className = 'hero-adversary glitch-out';
                    if (grain) grain.classList.remove('active');
                    if (flash) {
                        flash.classList.add('active');
                        setTimeout(function() { flash.classList.remove('active'); }, 200);
                    }
                    if (heroWrap) {
                        heroWrap.classList.add('shaking');
                        setTimeout(function() { heroWrap.classList.remove('shaking'); }, 300);
                    }
                    hero.style.transition = 'none';
                    hero.style.opacity = '0.5';
                    hero.style.filter = 'hue-rotate(-30deg) saturate(2) brightness(0.7)';
                    hero.style.transform = 'skewX(-3deg)';
                    setTimeout(function() {
                        hero.style.transition = 'all 0.6s ease';
                        hero.style.opacity = '';
                        hero.style.transform = '';
                        hero.style.filter = '';
                    }, 250);
                    setTimeout(function() {
                        adversary.className = 'hero-adversary';
                        adversaryActive = false;
                    }, 500);
                }, 2750);

                // Safety: guarantee unlock even if timers overlap
                setTimeout(function() {
                    if (adversaryActive) {
                        adversary.className = 'hero-adversary';
                        if (grain) grain.classList.remove('active');
                        hero.style.transition = 'all 0.3s ease';
                        hero.style.opacity = '';
                        hero.style.transform = '';
                        hero.style.filter = '';
                        adversaryActive = false;
                    }
                }, 4000);
            } else {
                if (ghost) {
                    ghost.textContent = symbols[Math.floor(Math.random() * symbols.length)];
                    ghost.classList.remove('visible');
                    ghost.offsetHeight;
                    ghost.classList.add('visible');
                }
                setTimeout(function() {
                    hero.style.transform = 'skewX(' + (-skew * 0.6) + 'deg) scale(0.98)';
                    hero.style.filter = 'hue-rotate(-30deg) saturate(2) brightness(0.8)';
                }, 400);
                setTimeout(function() {
                    hero.style.transform = 'skewX(' + (skew * 0.3) + 'deg)';
                    hero.style.filter = 'hue-rotate(20deg) brightness(1.1)';
                }, 900);
                setTimeout(function() {
                    hero.style.transition = 'all 0.5s ease';
                    hero.style.transform = '';
                    hero.style.filter = '';
                    if (ghost) ghost.classList.remove('visible');
                }, 1800);
            }
        }
        setInterval(function() {
            if (Math.random() < 0.5) heroGlitch();
        }, 5000);
    })();

    // Sweeping scanline every 20s
    (function() {
        var sweep = document.getElementById('sweep-scanline');
        function triggerSweep() {
            if (document.documentElement.getAttribute('data-animations') === 'off') return;
            sweep.style.animation = 'none';
            sweep.offsetHeight;
            sweep.style.animation = 'sweepDown 6s linear forwards';
        }
        setTimeout(triggerSweep, 2000);
        setInterval(triggerSweep, 20000);
    })();

    // Modal overlay for detection endpoints (faked streaming)
    var modalOverlay = document.getElementById('modal-overlay');
    var modalTitle = document.getElementById('modal-title');
    var modalBody = document.getElementById('modal-body');
    var typewriterTimer = null;

    function typewrite(el, text, speed, done) {
        var i = 0;
        function step() {
            if (i < text.length) {
                var chunk = text.slice(i, i + Math.ceil(Math.random() * 3));
                el.textContent += chunk;
                i += chunk.length;
                el.scrollTop = el.scrollHeight;
                typewriterTimer = setTimeout(step, speed + Math.random() * speed);
            } else if (done) {
                done();
            }
        }
        step();
    }

    var bootLines = [
        '[ system ] Initializing shell environment...',
        '[ system ] Connecting to container runtime...',
        '[ system ] Resolving execution context...',
        '[ exec   ] Running command...',
        ''
    ];

    function openModal(href, title) {
        modalTitle.textContent = '❯ ' + title;
        modalBody.textContent = '';
        modalOverlay.classList.add('open');

        var lineIdx = 0;
        var fetchDone = false;
        var fetchResult = null;
        var fetchErr = null;
        var bootDone = false;

        fetch(href)
            .then(function(res) { return res.text(); })
            .then(function(text) { fetchResult = text; fetchDone = true; showResultIfReady(); })
            .catch(function(err) { fetchErr = err; fetchDone = true; showResultIfReady(); });

        function showNextBootLine() {
            if (lineIdx < bootLines.length) {
                modalBody.textContent += bootLines[lineIdx] + '\n';
                modalBody.scrollTop = modalBody.scrollHeight;
                lineIdx++;
                typewriterTimer = setTimeout(showNextBootLine, 300 + Math.random() * 400);
            } else {
                bootDone = true;
                showResultIfReady();
            }
        }

        function showResultIfReady() {
            if (!bootDone || !fetchDone) return;
            if (fetchErr) {
                modalBody.textContent += '\n[ ERROR ] ' + fetchErr.message + '\n';
                return;
            }
            var output = fetchResult || '(no output)';
            modalBody.textContent += '[ done   ] Output received. Rendering...\n\n';
            typewrite(modalBody, output, 8, function() {
                modalBody.textContent += '\n\n[ system ] Command complete.';
                modalBody.scrollTop = modalBody.scrollHeight;
            });
        }

        showNextBootLine();
    }

    function closeModal() {
        modalOverlay.classList.remove('open');
        modalBody.textContent = '';
        if (typewriterTimer) { clearTimeout(typewriterTimer); typewriterTimer = null; }
    }

    modalOverlay.addEventListener('click', function(e) {
        if (e.target === modalOverlay) closeModal();
    });

    document.addEventListener('keydown', function(e) {
        if (e.key === 'Escape') closeModal();
    });

    // Make entire card open modal on click
    (function() {
        document.querySelectorAll('li').forEach(function(li) {
            li.addEventListener('click', function(e) {
                e.preventDefault();
                var a = li.querySelector('a');
                if (a) openModal(a.href, a.textContent);
            });
        });
    })();
    </script>
</body>
</html>
`
