import { defineStore } from 'pinia'
import { getProjects } from '../api/axios'

const profile = {
  name: 'Nanda Surya Diffa',
  headline: 'Web Developer | Android Developer | API & Middleware',
  location: 'Jakarta, Indonesia',
  email: 'nandasurya861@gmail.com',
  whatsappUrl: 'https://wa.me/6285719797174',
  linkedinUrl: 'https://www.linkedin.com/in/nanda-surya-194209163/',
  summary:
    'Junior Software Developer at Bank Artha Graha International with one year of professional experience in middleware integration using Fiorano and web development using Laravel. Experienced in Android and web feature enhancement, requirement analysis, and technical documentation such as FSD and internal testing documents.',
}

const quickStats = [
  { label: 'Professional Experience', value: '1+ Year' },
  { label: 'Current Role', value: 'Developer Staff' },
  { label: 'Primary Domain', value: 'API & Middleware' },
  { label: 'Graduation', value: 'Nov 2024' },
]

const experiences = [
  {
    role: 'Developer Staff (Full-Time)',
    company: 'PT Bank Artha Graha Internasional Tbk.',
    location: 'Menteng, Jakarta',
    period: 'Jan 2025 - Present',
    highlights: [
      'Developed and improved Fiorano middleware integration flows for cross-system data exchange.',
      'Delivered Laravel-based enhancements including new features, bug fixes, and maintenance.',
      'Collaborated with stakeholders to clarify requirements and coordinate implementation.',
      'Prepared Functional Specification Document (FSD) and internal testing documentation for handover.',
    ],
  },
  {
    role: 'Full Stack Website & Android Developer (Internship)',
    company: 'PT Federal Izumi Manufacturing (Member of Astra Otoparts)',
    location: 'Cileungsi, Bogor',
    period: 'Feb 2023 - Feb 2024',
    highlights: [
      'Improved Android app for maintenance department: UI/UX updates, new menus, and bug fixes using Java and PHP API.',
      'Improved preventive and breakdown maintenance process flows in the maintenance division.',
      'Developed CodeIgniter-based website for model change process digitization across related sections.',
    ],
  },
]

const education = {
  institution: 'Politeknik STMI Jakarta Cempaka Putih',
  major: 'Information System Automotive Industry',
  grade: 'GPA 3.67/4.00',
  period: 'Sep 2020 - Nov 2024',
}

const organizations = [
  {
    name: 'Information Systems Student Association',
    period: 'Jan 2021 - Jan 2023',
    notes: [
      'Chief Executive of Cadre Event.',
      'PSDM staff in BPH supervision area.',
    ],
  },
  {
    name: 'Student Representative Council',
    period: 'Aug 2023 - Sep 2023',
    notes: [
      'Head of Commission C for monitoring and evaluation of Student Executive Board performance.',
      'Participated in open session committee as chairman of event division.',
    ],
  },
]

const certifications = [
  {
    name: 'Development of Information System Business Process in Manufacturing Industry',
    issuer: 'BNSP',
    period: 'Jun 2024 - Jun 2027',
  },
]

const relatedCourses = [
  'Business and System Analyst',
  'Programming',
  'ERP',
  'Project Management Information System',
  'Information System Evaluation and Audit',
]

const skillGroups = {
  language: ['Indonesian', 'English (Intermediate)'],
  coreSkills: [
    'Web Development',
    'Mobile Android Development',
    'API and Middleware',
    'Business Analyst',
    'System Analyst',
    'Business Requirement Document',
    'Functional Specification Document',
    'Critical Thinking',
    'Analytical Thinking',
    'Problem Solving',
    'Public Speaking',
    'Leadership',
    'Teamwork',
  ],
  programmingLanguages: ['HTML5', 'CSS', 'JavaScript', 'TypeScript', 'Java', 'PHP', 'Python', 'Kotlin'],
  frameworks: ['Laravel', 'CodeIgniter', 'React.js', 'Vite.js', 'Node.js', 'Next.js', 'Vue.js', 'Tailwind CSS', 'Bootstrap'],
  tools: [
    'Visual Studio Code',
    'Android Studio',
    'Postman',
    'SoapUI',
    'Microsoft Office',
    'Bizagi',
    'Figma',
    'Canva',
    'Draw.io',
    'Fiorano',
    'WinSCP',
    'Mobaxterm',
  ],
  databases: ['MySQL', 'MSSQL', 'PostgreSQL', 'Firebase'],
  apiMiddleware: ['REST API', 'JSON', 'Fiorano'],
}

const featuredProjects = [
  {
    tag: 'Middleware Integration',
    title: 'Fiorano Flow Enhancement',
    company: 'PT Bank Artha Graha Internasional Tbk.',
    period: '2025 - Present',
    description:
      'Developed and enhanced middleware integration flows to support reliable data exchange across internal systems.',
    outcomes: [
      'Improved reliability of integration workflows for internal services.',
      'Reduced manual troubleshooting through continuous flow improvement.',
    ],
    stack: ['Fiorano', 'JSON', 'Middleware Integration'],
  },
  {
    tag: 'Web Enhancement',
    title: 'Laravel Business Feature Delivery',
    company: 'PT Bank Artha Graha Internasional Tbk.',
    period: '2025 - Present',
    description:
      'Implemented Laravel-based enhancements, bug fixes, and maintenance to align web applications with business requirements.',
    outcomes: [
      'Delivered production-ready feature updates with stakeholder coordination.',
      'Improved maintainability through clear requirement and testing documents.',
    ],
    stack: ['Laravel', 'PHP', 'MySQL'],
  },
  {
    tag: 'Android Development',
    title: 'Maintenance App UX and Feature Revamp',
    company: 'PT Federal Izumi Manufacturing',
    period: '2023 - 2024',
    description:
      'Upgraded maintenance Android app by redesigning UI/UX, adding new menu flows, and fixing defects in coordination with department needs.',
    outcomes: [
      'Supported smoother preventive and breakdown maintenance reporting.',
      'Improved usability for maintenance team workflows.',
    ],
    stack: ['Java', 'Android', 'PHP API'],
  },
  {
    tag: 'Process Digitization',
    title: 'Model Change Process Web System',
    company: 'PT Federal Izumi Manufacturing',
    period: '2023 - 2024',
    description:
      'Built a CodeIgniter-based website for model change process digitization involving Model Change and Precision Quality sections.',
    outcomes: [
      'Moved manual coordination into a structured digital workflow.',
      'Improved visibility and traceability of process status.',
    ],
    stack: ['CodeIgniter', 'PHP', 'MySQL'],
  },
]

export const usePortfolioStore = defineStore('portfolio', {
  state: () => ({
    profile,
    quickStats,
    experiences,
    education,
    organizations,
    certifications,
    relatedCourses,
    skillGroups,
    featuredProjects,
    remoteProjects: [],
    loadingRemoteProjects: false,
    remoteProjectsError: '',
    remoteProjectsLoaded: false,
  }),
  getters: {
    allSkills: (state) => {
      const merged = [
        ...state.skillGroups.coreSkills,
        ...state.skillGroups.programmingLanguages,
        ...state.skillGroups.frameworks,
        ...state.skillGroups.apiMiddleware,
      ]

      return [...new Set(merged)]
    },
  },
  actions: {
    async fetchRemoteProjects(force = false) {
      if (this.remoteProjectsLoaded && !force) {
        return this.remoteProjects
      }

      this.loadingRemoteProjects = true
      this.remoteProjectsError = ''

      try {
        const items = await getProjects()
        this.remoteProjects = Array.isArray(items) ? items : []
        this.remoteProjectsLoaded = true
      } catch {
        this.remoteProjects = []
        this.remoteProjectsLoaded = false
        this.remoteProjectsError = 'Unable to load additional projects from backend API.'
      } finally {
        this.loadingRemoteProjects = false
      }

      return this.remoteProjects
    },
  },
})
