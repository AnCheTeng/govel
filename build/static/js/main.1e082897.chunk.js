(this["webpackJsonpgovel-ui"]=this["webpackJsonpgovel-ui"]||[]).push([[0],{52:function(e,t,a){"use strict";a.r(t);var n=a(2),c=a(0),r=a.n(c),i=a(13),s=a.n(i),o=a(36),l=a(16),j=a(76),u=a(86),b=a(78),d=a(79),x=a(87),h=Object(j.a)((function(e){return{root:{maxWidth:"95%"},avatar:{marginTop:e.spacing(1.5)}}}));function O(){var e=h();return Object(n.jsxs)(b.a,{container:!0,item:!0,className:e.root,children:[Object(n.jsxs)(b.a,{container:!0,direction:"column",item:!0,xs:11,children:[Object(n.jsx)(d.a,{variant:"h4",color:"primary",children:"Flashcards"}),Object(n.jsx)(d.a,{variant:"button",color:"textPrimary",children:"Pick a set to practice"})]}),Object(n.jsx)(b.a,{container:!0,item:!0,xs:1,justify:"flex-end",children:Object(n.jsx)(x.a,{className:e.avatar,children:"AC"})})]})}var f=a(80),m=a(81),p=a(34),v=a.n(p),g=a(32),y=a.n(g),w=a(33),N=a.n(w),F=Object(j.a)((function(e){return{root:{flexGrow:1,transform:"rotate(-90deg)"}}}));function C(){var e=F(),t=r.a.useState("favorites"),a=Object(l.a)(t,2),c=a[0],i=a[1];return Object(n.jsxs)(f.a,{value:c,onChange:function(e,t){i(t)},className:e.root,children:[Object(n.jsx)(m.a,{label:"Favorites",value:"favorites",icon:Object(n.jsx)(y.a,{})}),Object(n.jsx)(m.a,{label:"Nearby",value:"nearby",icon:Object(n.jsx)(N.a,{})}),Object(n.jsx)(m.a,{label:"Folder",value:"folder",icon:Object(n.jsx)(v.a,{})})]})}var I=a(82),S=a(88),T=a(83),k=a(85),L=a(84),P=a(26),B=a.n(P),E=a(35),M=a.n(E),W=a(25),D=a.n(W),G=Object(j.a)((function(e){return{root:{display:"flex",justifyContent:"space-between"},details:{display:"flex",flexDirection:"column"},content:{flex:"1 0 auto"},hiddenText:{width:"10rem"},cover:{width:151},controls:{display:"flex",alignItems:"center",paddingLeft:e.spacing(1),paddingBottom:e.spacing(1)},playIcon:{height:38,width:38}}}));function J(e){var t=G(),a=Object(I.a)(),r=function(){var e=["nature","water","dog","cat","capybara","otter","beaver","tiger","wild"],t="https://source.unsplash.com/640x320/?"+e[Math.floor(Math.random()*e.length)],a=Object(c.useState)(t),n=Object(l.a)(a,2),r=n[0],i=n[1];return Object(c.useEffect)((function(){fetch(t,{redirect:"manual"}).then((function(e){return i(e.url)}))}),[]),r}();return Object(n.jsxs)(S.a,{className:t.root,children:[Object(n.jsxs)("div",{className:t.details,children:[Object(n.jsxs)(T.a,{className:t.content,children:[Object(n.jsx)(d.a,{component:"h5",variant:"h5",children:e.name}),Object(n.jsx)(d.a,{variant:"subtitle1",color:"textSecondary",noWrap:!0,className:t.hiddenText,children:e.chapter})]}),Object(n.jsxs)("div",{className:t.controls,children:[Object(n.jsx)(L.a,{"aria-label":"previous",children:"rtl"===a.direction?Object(n.jsx)(D.a,{}):Object(n.jsx)(B.a,{})}),Object(n.jsx)(L.a,{"aria-label":"play/pause",href:e.url,children:Object(n.jsx)(M.a,{className:t.playIcon})}),Object(n.jsx)(L.a,{"aria-label":"next",children:"rtl"===a.direction?Object(n.jsx)(B.a,{}):Object(n.jsx)(D.a,{})})]})]}),Object(n.jsx)(k.a,{className:t.cover,image:r,title:"Live from space album cover"})]})}var U=Object(j.a)((function(e){return{root:{flexGrow:1}}}));var A=function(){var e=U(),t=Object(c.useState)([{chapter:"---",name:"..........",url:"https://thawing-crag-57028.herokuapp.com/api/v1/lastUpdate"}]),a=Object(l.a)(t,2),r=a[0],i=a[1];return Object(c.useEffect)((function(){fetch("/api/v1/lastUpdate").then((function(e){return e.json()})).then((function(e){return i(e)}))}),[]),Object(n.jsx)(u.a,{maxWidth:"lg",children:Object(n.jsxs)(b.a,{container:!0,className:e.root,spacing:3,direction:"column",children:[Object(n.jsx)(b.a,{container:!0,item:!0,justify:"center",children:Object(n.jsx)(O,{})}),Object(n.jsxs)(b.a,{container:!0,item:!0,children:[Object(n.jsx)(b.a,{container:!0,item:!0,xs:2,wrap:"wrap",justify:"center",alignItems:"center",children:Object(n.jsx)(C,{})}),Object(n.jsx)(b.a,{container:!0,item:!0,xs:!0,justify:"center",spacing:2,direction:"column",children:r.map((function(e){return Object(n.jsx)(b.a,{item:!0,children:Object(n.jsx)(J,Object(o.a)({},e))},e.name)}))})]})]})})},q=function(e){e&&e instanceof Function&&a.e(3).then(a.bind(null,89)).then((function(t){var a=t.getCLS,n=t.getFID,c=t.getFCP,r=t.getLCP,i=t.getTTFB;a(e),n(e),c(e),r(e),i(e)}))};s.a.render(Object(n.jsx)(r.a.StrictMode,{children:Object(n.jsx)(A,{})}),document.getElementById("root")),q()}},[[52,1,2]]]);
//# sourceMappingURL=main.1e082897.chunk.js.map