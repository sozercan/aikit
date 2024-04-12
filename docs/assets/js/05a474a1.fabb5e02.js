"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[73],{732:(e,t,s)=>{s.r(t),s.d(t,{assets:()=>l,contentTitle:()=>a,default:()=>u,frontMatter:()=>r,metadata:()=>o,toc:()=>c});var n=s(5893),i=s(1151);const r={title:"Release Process"},a=void 0,o={id:"release",title:"Release Process",description:"The release process is as follows:",source:"@site/docs/release.md",sourceDirName:".",slug:"/release",permalink:"/aikit/release",draft:!1,unlisted:!1,editUrl:"https://github.com/sozercan/aikit/blob/main/website/docs/docs/release.md",tags:[],version:"current",frontMatter:{title:"Release Process"},sidebar:"sidebar",previous:{title:"Stable Diffusion",permalink:"/aikit/stablediffusion"}},l={},c=[];function d(e){const t={a:"a",admonition:"admonition",code:"code",li:"li",p:"p",pre:"pre",ul:"ul",...(0,i.a)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(t.p,{children:"The release process is as follows:"}),"\n",(0,n.jsxs)(t.ul,{children:["\n",(0,n.jsxs)(t.li,{children:["Tag the ",(0,n.jsx)(t.code,{children:"main"})," branch with a version number that's semver compliant (vMAJOR.MINOR.PATCH), and push the tag to GitHub."]}),"\n"]}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-bash",children:"git tag v0.1.0\ngit push origin v0.1.0\n"})}),"\n",(0,n.jsxs)(t.ul,{children:["\n",(0,n.jsxs)(t.li,{children:["\n",(0,n.jsxs)(t.p,{children:["GitHub Actions will automatically build the AIKit image and push the versioned and ",(0,n.jsx)(t.code,{children:"latest"})," tag to GitHub Container Registry (GHCR) using ",(0,n.jsx)(t.a,{href:"https://github.com/sozercan/aikit/actions/workflows/release.yaml",children:"release action"}),"."]}),"\n"]}),"\n",(0,n.jsxs)(t.li,{children:["\n",(0,n.jsxs)(t.p,{children:["Once release is done, trigger ",(0,n.jsx)(t.a,{href:"https://github.com/sozercan/aikit/actions/workflows/update-models.yaml",children:"update models"})," action to update the pre-built models."]}),"\n"]}),"\n"]}),"\n",(0,n.jsxs)(t.admonition,{type:"note",children:[(0,n.jsx)(t.p,{children:"At this time, Mixtral 8x7b model does not fit into GitHub runners due to its size. It is built and pushed to GHCR manually."}),(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-shell",children:"docker buildx build . -t ghcr.io/sozercan/mixtral:8x7b \\\n  -t ghcr.io/sozercan/mixtral:8x7b-instruct \\\n  -f models/mixtral-7x8b-instruct.yaml \\\n  --push --progress=plain --provenance=true --sbom=true\n"})}),(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-shell",children:"docker buildx build . -t ghcr.io/sozercan/mixtral:8x7b-cuda -t ghcr.io/sozercan/mixtral:8x7b-instruct-cuda \\\n  -f models/mixtral-7x8b-instruct-cuda.yaml \\\n  --push --progress=plain --provenance=true --sbom=true\n"})})]})]})}function u(e={}){const{wrapper:t}={...(0,i.a)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(d,{...e})}):d(e)}},1151:(e,t,s)=>{s.d(t,{Z:()=>o,a:()=>a});var n=s(7294);const i={},r=n.createContext(i);function a(e){const t=n.useContext(r);return n.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function o(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(i):e.components||i:a(e.components),n.createElement(r.Provider,{value:t},e.children)}}}]);