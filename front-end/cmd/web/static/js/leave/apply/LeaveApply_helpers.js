class LeaveApply_Helpers {

    // Adding file names into form so database can register
    populateInputHidden(id, value) {
        const element = document.getElementById(id);
        if (element.tagName.toLowerCase() === "input") {
            element.value = value;
        }
    }   
}