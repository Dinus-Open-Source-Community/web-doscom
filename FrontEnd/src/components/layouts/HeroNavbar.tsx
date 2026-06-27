import { useState, useEffect } from 'react'
import { Close, Join } from '../../assets/icons'
import { ROUTES } from '../../lib/routes'

interface Props {
  logoHtml: string
  currentPath: string
}

const navItems = [
  { name: 'Home', href: ROUTES.home },
  { name: 'About', href: ROUTES.about },
  { name: 'Division', href: ROUTES.division.list },
  { name: 'Works', href: ROUTES.works },
  { name: 'Blog', href: ROUTES.blog },
  { name: 'Contact', href: ROUTES.contact },
]

export default function Navbar({ logoHtml, currentPath }: Props) {
  const [isOpen, setIsOpen] = useState(false)

  useEffect(() => {
    document.body.style.overflow = isOpen ? 'hidden' : 'unset'

    const visibilityEvent = new CustomEvent('toggle-sticky-visibility', {
      detail: { hidden: isOpen },
    })
    window.dispatchEvent(visibilityEvent)

    const handleOpenEvent = () => setIsOpen(true)
    window.addEventListener('open-mobile-menu', handleOpenEvent)

    return () => {
      window.removeEventListener('open-mobile-menu', handleOpenEvent)
    }
  }, [isOpen])

  return (
    <>
      <nav className="absolute top-8 left-0 z-50 flex w-full justify-center px-4">
        <div className="flex w-[92vw] max-w-[1500px] items-center justify-between rounded-full bg-white px-8 py-3 text-black shadow-xl backdrop-blur-md 2xl:w-[85vw]">
          <div className="flex items-center">
            <a
              href={ROUTES.home}
              className="w-32 text-[#0A84DC] md:w-36 [&>svg]:h-full [&>svg]:w-full"
              dangerouslySetInnerHTML={{ __html: logoHtml }}
            />
          </div>

          <ul className="hidden gap-8 text-lg font-normal lg:flex">
            {navItems.map(item => (
              <li key={item.href}>
                <a
                  href={item.href}
                  className={`transition-colors duration-200 hover:text-[#0A84DC] ${
                    currentPath === item.href
                      ? 'font-semibold text-[#0A84DC]'
                      : ''
                  }`}
                >
                  {item.name}
                </a>
              </li>
            ))}
          </ul>

          <div className="flex items-center gap-3">
            <button className="flex items-center justify-center rounded-full bg-[#0A84DC] px-4 py-2 font-semibold text-white transition-transform hover:scale-105 active:scale-95 md:px-6 md:py-2">
              <span className="hidden md:block">Join Us</span>
              <div
                className="h-5 w-5 md:hidden [&>svg]:h-full [&>svg]:w-full"
                dangerouslySetInnerHTML={{ __html: Join }}
              />
            </button>

            <button
              onClick={() => setIsOpen(true)}
              className="flex h-10 w-10 flex-col items-center justify-center gap-1 rounded-full bg-gray-100 transition-colors hover:bg-gray-200 lg:hidden"
            >
              <span className="h-0.5 w-4 bg-black"></span>
              <span className="h-0.5 w-4 bg-black"></span>
            </button>
          </div>
        </div>
      </nav>

      <div
        className={`fixed inset-0 z-[100] flex flex-col bg-[#F5F5F3] p-8 transition-all duration-500 ease-in-out ${
          isOpen ? 'translate-y-0 opacity-100' : '-translate-y-full opacity-0'
        }`}
      >
        <div className="flex items-center justify-between">
          <div
            className="w-32 opacity-30"
            dangerouslySetInnerHTML={{ __html: logoHtml }}
          />

          <button
            onClick={() => setIsOpen(false)}
            className="flex h-12 w-12 items-center justify-center rounded-full bg-black/30 text-[#F5F5F3] backdrop-blur-sm transition-colors hover:bg-black/50"
          >
            <span
              className="h-6 w-6 [&>svg]:h-full [&>svg]:w-full"
              dangerouslySetInnerHTML={{ __html: Close }}
            />
          </button>
        </div>

        <div className="flex flex-1 flex-col items-center justify-center gap-6">
          {navItems.map((item, index) => (
            <a
              key={item.href}
              href={item.href}
              style={{
                transitionDelay: isOpen ? `${index * 200}ms` : '0ms',
              }}
              className={`transform text-3xl font-bold tracking-tight transition-all duration-500 ${
                isOpen
                  ? 'translate-y-0 opacity-100'
                  : 'translate-y-10 opacity-0'
              } hover:text-[#0A84DC] ${
                currentPath === item.href ? 'text-[#0A84DC]' : 'text-black'
              }`}
            >
              {item.name}
            </a>
          ))}
        </div>

        <div
          style={{
            // Delay = (jumlah item * 200ms) + sedikit buffer biar pas
            transitionDelay: isOpen ? `${navItems.length * 150}ms` : '0ms',
          }}
          className={`flex transform justify-between border-t border-gray-200 pt-8 text-sm font-medium text-gray-500 transition-all duration-500 ${
            isOpen ? 'translate-y-0 opacity-100' : 'translate-y-5 opacity-0'
          }`}
        >
          <div className="flex gap-6">
            <a href="https://instagram.com" className="hover:text-black">
              Instagram
            </a>
            <a href="https://github.com" className="hover:text-black">
              Github
            </a>
          </div>
          <span>DOSCOM</span>
        </div>
      </div>
    </>
  )
}
