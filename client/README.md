# Client: Built using SvelteKit

## Developing

Once you've created a project and installed dependencies with `npm install` (or `pnpm install` or `yarn`), start a development server:

```bash
npm run dev

# or start the server and open the app in a new browser tab
npm run dev -- --open

# or start the server and open to local network
npm run dev -- --host
```

## Building

Before creating a production version of your app, install an [adapter](https://kit.svelte.dev/docs#adapters) for your target environment. Then:

```bash
npm run build
```

> You can preview the built app with `npm run preview`, regardless of whether you installed an adapter. This should _not_ be used to serve your app in production.

## File structure

```txt
src
├── app.html -- loaded by browser, injected with svelte app
├── components -- standalone, reusable components (ie. buttons, input, etc.)
├── global.d.ts -- something to do with sveltekit
├── routes -- browser routes (ie '/home'), index = '/'
│   └── index.svelte -- home page, renders views 
├── scripts -- TypeScript helper files
│   ├── appState.ts -- handles states that should be globally accessible
│   └── ws.ts -- handles all things websocket related
└── views -- individual views to be rendered on screen
    ├── home.svelte       
    ├── table.svelte      
    └── waitingroom.svelte
```