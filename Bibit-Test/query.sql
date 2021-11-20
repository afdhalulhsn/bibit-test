select u2.id ,u2.username,u.username as parent from "user" u right join "user" u2
on u2.parent = u.id;