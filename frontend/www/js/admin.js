
$(function() {
    var experience = document.getElementById("experience-timeline");
    var education =  document.getElementById("education-list");
    var projects =   document.getElementById("project-list");
    var skills =     document.getElementById("skill-list");
    var social =     document.getElementById("social-link-list");

    $("#add-job").on('click', function() {
        var newJob = experience.children[0].cloneNode(true);
        experience.appendChild(newJob);
    });

    $("#remove-job").on('click', function() {
        if (experience.children.length > 1) {
            var lastChild = experience.children[experience.children.length - 1];
            experience.removeChild(lastChild);
        }
    });

    $("#add-school").on('click', function() {
        var newSchool = education.children[0].cloneNode(true);
        education.appendChild(newSchool);
    });

    $("#remove-school").on('click', function() {
       if (education.children.length > 1) {
           var lastChild = education.children[education.children.length - 1];
           education.removeChild(lastChild);
       }
    });

    $("#add-project").on('click', function() {
        var newProject = projects.children[0].cloneNode(true);
        projects.appendChild(newProject);
    });

    $("#remove-project").on('click', function() {
        if (projects.children.length > 1) {
            var lastChild = projects.children[projects.children.length - 1];
            projects.removeChild(lastChild);
        }
    });

    $("#add-skill").on('click', function() {
        var newSkill = skills.children[0].cloneNode(true);
        skills.appendChild(newSkill);
    });

    $("#remove-skill").on('click', function() {
        if (skills.children.length > 1) {
            var lastChild = skills.children[skills.children.length - 1];
            skills.removeChild(lastChild);
        }
    });

    $("#add-social").on('click', function() {
        var newSocial = social.children[0].cloneNode(true);
        social.appendChild(newSocial);
    });

    $("#remove-social").on('click', function() {
        if (social.children.length > 1) {
            var lastChild = social.children[social.children.length - 1];
            social.removeChild(lastChild);
        }
    });

    // Change handler for profile picture
    $("#profile-pic-url").on('change', function() {
       $("#profile-picture").attr('src', $("#profile-pic-url").val());
    });

    // Save button
    $("#save-button").on('click', function() {
        var data = {};

        data["About"] = $("#about-text").val();
        data["ProfilePicURL"] = $("#profile-pic-url").val();
        data["ResumeURL"] = $("#resume-url").val();

        data["Experience"] = getExperience();
        data["Education"] = getEducation();
        data["Projects"] = getProjects();
        data["Skills"] = [];

        $(".skill-item").each(function() {
           data["Skills"].push($("input", this).val());
        });

        data["SocialLinks"] = [];

        $(".social-link-entry").each(function () {
            data["SocialLinks"].push({
                Link: $("input[name=icon-link]", this).val(),
                IconClass: $("input[name=icon-class]", this).val()
            });
        });

        console.log(data);

        $.ajax({
            url: '/admin/save',
            type: 'post',
            dataType: 'json',
            success: function (data) {
                $('#target').html(data.msg);
            },
            data: JSON.stringify(data)
        });

    });

});

function getExperience() {
    var experience = [];

    $(".job-data").each(function() {
        var currentExp = {};
        currentExp["Timeline"] = $("input[name=timeline]", this).val();
        currentExp["Company"] = $("input[name=company]", this).val();
        currentExp["Title"] = $("input[name=title]", this).val();
        currentExp["Description"] = $("textarea", this).val();
        experience.push(currentExp);
    });

    return experience;
}

function getEducation() {
    var education = [];

    $(".education-block").each(function() {
        var currentEdu = {};
        currentEdu["Timeline"] = $("input[name=timeline]", this).val();
        currentEdu["Name"] = $("input[name=name]", this).val();
        currentEdu["Major"] = $("input[name=major]", this).val();
        currentEdu["DescriptionList"] = $("textarea", this).val().split("\n").filter(function(el) {
            return el.length > 0;
        });
        education.push(currentEdu);
    });

    return education;
}

function getProjects() {
    var projects = [];

    $(".project-info").each(function () {
        var currentProj = {};
        currentProj["Year"] = parseInt($("input[name=year]", this).val());
        currentProj["Name"] = $("input[name=name]", this).val();
        currentProj["LinkText"] = $("input[name=link-text]", this).val();
        currentProj["LinkURL"] = $("input[name=link-url]", this).val();
        currentProj["ImageURL"] = $("input[name=image-url]", this).val();
        currentProj["Description"] = $("textarea", this).val();
        projects.push(currentProj);
    });

    return projects;
}
