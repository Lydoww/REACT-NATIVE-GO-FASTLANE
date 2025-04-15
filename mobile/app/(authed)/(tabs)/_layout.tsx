import { TabBarIcon } from "@/components/navigation/TarBarIcon";
import Ionicons from "@expo/vector-icons/Ionicons";
import { Tabs } from "expo-router";
import { ComponentProps } from "react";
import { Text } from "react-native";

export default function TabLayout() {
  const tabs = [
    {
      showFor: [],
      name: "(events)",
      displayName: "Events",
      icon: "calendar",
      options: {
        headerShown: false,
      },
    },
    {
      showFor: [],
      name: "(tickets)",
      displayName: "My Tickets",
      icon: "ticket",
      options: {
        headerShown: false,
      },
    },
    {
      showFor: [],
      name: "(scan-ticket)",
      displayName: "Scan Ticket",
      icon: "scan",
      options: {
        headerShown: true,
      },
    },
    {
      showFor: [],
      name: "(settings)",
      displayName: "Settings",
      icon: "cog",
      options: {
        headerShown: true,
      },
    },
  ];

  return (
    <Tabs>
      {tabs.map((tab) => {
        return (
          <Tabs.Screen
            key={tab.name}
            name={tab.name}
            options={{
              ...tab.options,
              headerTitle: tab.displayName,
              tabBarLabel: ({ focused }) => (
                <Text
                  style={{ color: focused ? "black" : "gray", fontSize: 12 }}
                >
                  {tab.displayName}
                </Text>
              ),
              tabBarIcon: ({ focused }) => {
                return (
                  <Ionicons
                    name={tab.icon as ComponentProps<typeof Ionicons>["name"]}
                    size={24}
                    color={focused ? "black" : "gray"}
                  />
                );
              },
            }}
          />
        );
      })}
    </Tabs>
  );
}
