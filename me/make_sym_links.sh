#!/bin/sh

ln -sfh ../common/main.mustache              en/template.mustache
ln -sfh ../../common/header.mustache         en/header/template.mustache
ln -sfh ../../common/education.mustache      en/education/template.mustache
ln -sfh ../../common/pro_experience.mustache en/professional/template.mustache
ln -sfh ../../common/skills.mustache         en/skills/template.mustache
ln -sfh ../../common/hobby.mustache          en/hobby/template.mustache

ln -sfh ../../en/resume.json       web/en/resume.json
ln -sfh ../../en/education         web/en/education
ln -sfh ../../en/professional      web/en/professional
ln -sfh ../../en/skills            web/en/skills
ln -sfh ../../en/hobby             web/en/hobby

