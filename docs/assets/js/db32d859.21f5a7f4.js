"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[884],{350:(e,t,i)=>{i.r(t),i.d(t,{assets:()=>c,contentTitle:()=>r,default:()=>h,frontMatter:()=>a,metadata:()=>o,toc:()=>l});var n=i(4848),s=i(8453);const a={title:"Supply Chain Security"},r=void 0,o={id:"security",title:"Supply Chain Security",description:"AIKit is designed with security in mind. Our approach to supply chain security includes detailed tracking of software components, transparent build processes, and proactive vulnerability management. This ensures that every part of our software ecosystem remains secure and trustworthy.",source:"@site/docs/security.md",sourceDirName:".",slug:"/security",permalink:"/aikit/docs/security",draft:!1,unlisted:!1,editUrl:"https://github.com/sozercan/aikit/blob/main/website/docs/docs/security.md",tags:[],version:"current",frontMatter:{title:"Supply Chain Security"},sidebar:"sidebar",previous:{title:"Kubernetes Deployment",permalink:"/aikit/docs/kubernetes"},next:{title:"Inference API Specifications",permalink:"/aikit/docs/specs-inference"}},c={},l=[{value:"SBOM (Software Bill of Materials)",id:"sbom-software-bill-of-materials",level:2},{value:"Provenance attestation",id:"provenance-attestation",level:2},{value:"Vulnerability Patching",id:"vulnerability-patching",level:2},{value:"Image Signature Verification",id:"image-signature-verification",level:2}];function d(e){const t={a:"a",code:"code",h2:"h2",li:"li",p:"p",pre:"pre",ul:"ul",...(0,s.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(t.p,{children:"AIKit is designed with security in mind. Our approach to supply chain security includes detailed tracking of software components, transparent build processes, and proactive vulnerability management. This ensures that every part of our software ecosystem remains secure and trustworthy."}),"\n",(0,n.jsx)(t.h2,{id:"sbom-software-bill-of-materials",children:"SBOM (Software Bill of Materials)"}),"\n",(0,n.jsxs)(t.p,{children:["AIKit publishes ",(0,n.jsx)(t.a,{href:"https://www.cisa.gov/sbom",children:"Software Bill of Materials (SBOM)"})," for each release and for all ",(0,n.jsx)(t.a,{href:"/aikit/docs/premade-models",children:"pre-made models"}),". The SBOM is a comprehensive list of all the components and dependencies used in the project, detailing their versions, licenses, and sources. This transparency helps users and stakeholders understand what software is included, facilitating better risk management and compliance with security and licensing requirements."]}),"\n",(0,n.jsx)(t.p,{children:"To access the SBOM for a specific AIKit image, use the following command:"}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-bash",children:'# update this with the image you want to inspect\nIMAGE=ghcr.io/sozercan/llama3:8b\ndocker buildx imagetools inspect $IMAGE --format "{{ json .SBOM.SPDX }}"\n'})}),"\n",(0,n.jsxs)(t.p,{children:["The output will provide a detailed JSON document listing all the software components in the image, including direct and transitive dependencies. For more information, please visit ",(0,n.jsx)(t.a,{href:"https://docs.docker.com/build/attestations/sbom/",children:"Docker SBOM documentation"}),"."]}),"\n",(0,n.jsx)(t.h2,{id:"provenance-attestation",children:"Provenance attestation"}),"\n",(0,n.jsx)(t.p,{children:"Provenance attestation provides a detailed record of how and where an image was built, offering transparency and trust in the build process. AIKit uses BuildKit to generate and publish provenance data for each of its images. This data includes information about the build environment, the build process, and the source control context, ensuring that the images are traceable and verifiable from their origins to their final state."}),"\n",(0,n.jsx)(t.p,{children:"To inspect the provenance attestation for an AIKit image, you can use the following command:"}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-bash",children:'# update this with the image you want to inspect\nIMAGE=ghcr.io/sozercan/llama3:8b\ndocker buildx imagetools inspect $IMAGE --format "{{ json .Provenance.SLSA }}"\n'})}),"\n",(0,n.jsxs)(t.p,{children:["This command will output a JSON file containing the build provenance details, including the source repository, commit hash, build configuration, and more. This helps verify that the image was built from trusted sources and has not been tampered with. For more information, please visit ",(0,n.jsx)(t.a,{href:"https://docs.docker.com/build/attestations/slsa-provenance/",children:"Docker Provenance documentation"}),"."]}),"\n",(0,n.jsx)(t.h2,{id:"vulnerability-patching",children:"Vulnerability Patching"}),"\n",(0,n.jsxs)(t.p,{children:["Ensuring that our images are free from known vulnerabilities is crucial. Not only AIKit uses a custom distroless-based base image to reduce the number of vulnerabilities, attack surface and size, AIKit uses ",(0,n.jsx)(t.a,{href:"https://github.com/project-copacetic/copacetic",children:"Copacetic"})," to scan and patch OS-based vulnerabilities for all ",(0,n.jsx)(t.a,{href:"/aikit/docs/premade-models",children:"pre-made models"})," on a weekly basis. Copacetic automates the process of identifying and remediating security issues, helping us maintain a robust and secure software supply chain."]}),"\n",(0,n.jsx)(t.p,{children:"Every week, Copacetic performs the following actions:"}),"\n",(0,n.jsxs)(t.ul,{children:["\n",(0,n.jsxs)(t.li,{children:["Scan: It analyzes the images for vulnerabilities using ",(0,n.jsx)(t.a,{href:"https://github.com/aquasecurity/trivy",children:"Trivy"})," against a comprehensive database of known security issues."]}),"\n",(0,n.jsxs)(t.li,{children:["Patch: It automatically applies patches or updates to mitigate any identified vulnerabilities using ",(0,n.jsx)(t.a,{href:"https://github.com/project-copacetic/copacetic",children:"Copacetic"}),"."]}),"\n",(0,n.jsx)(t.li,{children:"Publish: It updates the images with the latest security fixes and publishes them to our container registry."}),"\n"]}),"\n",(0,n.jsx)(t.p,{children:"This automated and regular process ensures that our users always have access to the most secure and up-to-date images. You can monitor the status and results of these scans on our security dashboard."}),"\n",(0,n.jsx)(t.h2,{id:"image-signature-verification",children:"Image Signature Verification"}),"\n",(0,n.jsxs)(t.p,{children:["AIKit and pre-made models are keyless signed with OIDC in GitHub Actions with ",(0,n.jsx)(t.a,{href:"https://github.com/sigstore/cosign",children:"cosign"}),". You can verify the images with the following commands:"]}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-bash",children:"IMAGE=ghcr.io/sozercan/llama2:7b # update this with the image you want to verify\nDIGEST=$(cosign triangulate ${IMAGE} --type digest)\ncosign verify ${DIGEST} \\\n    --certificate-oidc-issuer https://token.actions.githubusercontent.com \\\n    --certificate-identity-regexp 'https://github\\.com/sozercan/aikit/\\.github/workflows/.+'\n"})}),"\n",(0,n.jsx)(t.p,{children:"You should see an output similar to the following:"}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-bash",children:"Verification for ghcr.io/sozercan/llama2@sha256:d47fdba491a9a47ce4911539a77e0c0a12b2e14f5beed88cb8072924b02130b4 --\nThe following checks were performed on each of these signatures:\n  - The cosign claims were validated\n  - Existence of the claims in the transparency log was verified offline\n  - The code-signing certificate was verified using trusted certificate authority certificates\n...\n"})})]})}function h(e={}){const{wrapper:t}={...(0,s.R)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(d,{...e})}):d(e)}},8453:(e,t,i)=>{i.d(t,{R:()=>r,x:()=>o});var n=i(6540);const s={},a=n.createContext(s);function r(e){const t=n.useContext(a);return n.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function o(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(s):e.components||s:r(e.components),n.createElement(a.Provider,{value:t},e.children)}}}]);