let navIsOpen = false
const openNav = () => {
    let navBtn = document.getElementById("hamburger-nav")
    let hamBtn = document.querySelector("hamburger-btn")
    if(!navIsOpen) {
        navBtn.style.display = "flex"
        navIsOpen = true
    } else {
        navBtn.style.display = "none"
        navIsOpen = false
    }

    hamBtn = navIsOpen ? 'fa-solid fa-bars' :
   "fa-regular fa-xmark" 
}