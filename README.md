# Aditya Dwi Nugroho - Portfolio

> Retro-modern portfolio built with Go serverless functions + HTMX. No JavaScript frameworks. Pure vibes.

## 🎨 Design Philosophy

- **Retro Y2K Aesthetic**: Neon gradients, glass morphism, pixel fonts
- **Modern UX**: Smooth scrolling, no terminal commands, user-friendly
- **Minimal JS**: Only HTMX for dynamic content, smooth scroll for navigation
- **Full Stack**: Go backend, semantic HTML, custom CSS

## 🛠️ Tech Stack

### Backend
- **Go 1.21+**: Serverless functions
- **Vercel**: Deployment platform

### Frontend
- **HTML5**: Semantic markup
- **CSS3**: Custom styling with variables
- **HTMX**: Hypermedia-driven interactions

### Fonts
- **Press Start 2P**: Retro pixel headings
- **Space Mono**: Modern monospace body

## 📁 Project Structure

```
portfolio/
├── api/                    # Go serverless functions
│   ├── visitor.go         # Visitor counter
│   ├── projects.go        # Projects showcase
│   └── contact.go         # Contact form handler
├── public/                # Static files
│   ├── index.html        # Main page
│   └── styles.css        # All styling
├── go.mod                # Go module
└── vercel.json          # Vercel config
```

## 🚀 Features

### ✅ Implemented
- [x] Hero section with glitch effect
- [x] About cards with glass morphism
- [x] Skills grid with tech stack
- [x] Dynamic projects loading (HTMX)
- [x] Contact form with validation
- [x] Visitor counter
- [x] Responsive design
- [x] Smooth scrolling navigation
- [x] Retro color themes (magenta/cyan/yellow)

### 🎯 Customization Needed
1. **Personal Info**: Update name, links, email in `index.html`
2. **Projects**: Modify projects array in `api/projects.go`
3. **Skills**: Edit skill chips in `index.html`
4. **Social Links**: Update GitHub, LinkedIn, Email URLs
5. **Email Integration**: Add SMTP config in `api/contact.go`

## 🔧 Local Development

### Prerequisites
- Go 1.21 or higher
- Vercel CLI (optional)

### Run Locally

1. **Install Vercel CLI** (if not installed):
```bash
npm install -g vercel
```

2. **Run dev server**:
```bash
vercel dev
```

3. **Open browser**:
```
http://localhost:3000
```

## 📦 Deployment

### Deploy to Vercel

1. **Install Vercel CLI**:
```bash
npm install -g vercel
```

2. **Login**:
```bash
vercel login
```

3. **Deploy**:
```bash
vercel
```

4. **Production deployment**:
```bash
vercel --prod
```

### Environment Variables

No environment variables needed for basic setup. For production:

- `SMTP_HOST`: Email server (for contact form)
- `SMTP_PORT`: Email port
- `SMTP_USER`: Email username
- `SMTP_PASS`: Email password
- `TO_EMAIL`: Your email address

## 🎨 Color Customization

Edit CSS variables in `styles.css`:

```css
:root {
    --primary: #ff00ff;      /* Magenta */
    --secondary: #00ffff;    /* Cyan */
    --accent: #ffff00;       /* Yellow */
    --bg-dark: #0a0014;      /* Dark purple */
    --bg-mid: #1a0a2e;       /* Mid purple */
}
```

## 📝 Content Update Guide

### Update Projects
Edit `api/projects.go`:
```go
projects := []Project{
    {
        Title:       "Your Project Name",
        Description: "Your description",
        Tech:        []string{"Go", "React", "etc"},
        Link:        "https://github.com/...",
        Status:      "Production",
    },
}
```

### Update Skills
Edit `index.html` skill sections:
```html
<div class="skill-chip">YourSkill</div>
```

### Update About
Edit `index.html` about cards content.

## 🔒 Security Notes

- Form validation included (frontend + backend)
- Email regex validation
- CORS headers configured
- Input sanitization needed for production
- Rate limiting recommended for contact form

## 📊 Performance

- **No JavaScript frameworks**: Faster load times
- **Serverless functions**: Auto-scaling
- **Minimal CSS**: < 10KB
- **HTMX**: < 15KB
- **Total JS**: ~15KB (only HTMX)

## 🌐 Browser Support

- Chrome/Edge (latest)
- Firefox (latest)
- Safari (latest)
- Mobile browsers

## 📄 License

MIT License - Feel free to use and customize!

## 🤝 Contributing

This is a personal portfolio, but feel free to fork and create your own version!

## 📧 Contact

- Email: aditya@example.com (update this!)
- GitHub: [@adityadwi](https://github.com/adityadwi) (update this!)
- LinkedIn: [Aditya Nugroho](https://linkedin.com/in/adityanugroho) (update this!)

---

Built with ❤️ using Go + HTMX | No JavaScript frameworks were harmed
