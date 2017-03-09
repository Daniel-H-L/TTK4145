#/bin/bash

go build main.go

scp -r /Home/student/Desktop/gr9/Elevator student@129.241.187.158:gruppe9
#scp -r /home/student/Desktop/Sanntid/Elevator_project_gr_16/main student@129.241.187.161:gruppe16/main


gnome-terminal --title "virtual_2: server" -x ssh student@129.241.187.152 &
#gnome-terminal --title "virtual_3: server" -x ssh student@129.241.187.161 &

scp -r /home/student/Documents/Group18 student@129.241.187.152:test18


# Remove directory
rm -rf lampp